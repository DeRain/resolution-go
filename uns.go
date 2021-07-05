package resolution

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethclient"
	kns "github.com/jgimeno/go-namehash"
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
	"github.com/unstoppabledomains/resolution-go/uns/contracts/proxyreader"
	"github.com/unstoppabledomains/resolution-go/uns/contracts/registry"
)

// Uns is a naming service handles Unstoppable domains resolution.
type Uns struct {
	proxyReader     *proxyreader.Contract
	supportedKeys   supportedKeys
	contractBackend bind.ContractBackend
	metadataClient  MetadataClient
}

// UnsBuilder is a builder to setup and build instance of Uns service.
type UnsBuilder interface {
	// SetContractBackend set Ethereum backend for communication with CNS registry
	SetContractBackend(backend bind.ContractBackend) UnsBuilder

	// SetMetadataClient set http backend for communication with ERC721 metadata server
	SetMetadataClient(backend MetadataClient) UnsBuilder

	// Build Uns instance
	Build() (*Uns, error)
}

type unsBuilder struct {
	contractBackend bind.ContractBackend
	metadataClient  MetadataClient
}

type MetadataClient interface {
	Get(url string) (resp *http.Response, err error)
}

const unsProvider = "https://rinkeby.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e"
const cnsEventsStartingBlock uint64 = 9923764
const unsEventsStartingBlock uint64 = 99999999

var unsZeroAddress = common.HexToAddress("0x0")
var unsMainnetProxyReader = common.HexToAddress("0x299974AeD8911bcbd2C61262605b89F591a53E83")
var cnsMainnetDefaultResolver = common.HexToAddress("0xb66DcE2DA6afAAa98F2013446dBCB0f4B0ab2842")

// NewUnsBuilder Creates builder to setup new instance of Uns service.
func NewUnsBuilder() UnsBuilder {
	return &unsBuilder{}
}

// SetContractBackend set Ethereum backend for communication with UNS registry
func (cb *unsBuilder) SetContractBackend(backend bind.ContractBackend) UnsBuilder {
	cb.contractBackend = backend
	return cb
}

func (cb *unsBuilder) SetMetadataClient(client MetadataClient) UnsBuilder {
	cb.metadataClient = client
	return cb
}

// Build Uns instance
func (cb *unsBuilder) Build() (*Uns, error) {
	if cb.contractBackend == nil {
		backend, err := ethclient.Dial(unsProvider)
		if err != nil {
			return nil, err
		}
		cb.contractBackend = backend
	}
	if cb.metadataClient == nil {
		cb.metadataClient = &http.Client{}
	}
	proxyReaderContract, err := proxyreader.NewContract(unsMainnetProxyReader, cb.contractBackend)
	if err != nil {
		return nil, err
	}
	supportedKeys, err := newSupportedKeys()
	if err != nil {
		return nil, err
	}
	return &Uns{proxyReader: proxyReaderContract, supportedKeys: supportedKeys, contractBackend: cb.contractBackend, metadataClient: cb.metadataClient}, nil
}

// Data Get raw data attached to domain
func (c *Uns) Data(domainName string, keys []string) (*struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	normalizedName := normalizeName(domainName)
	if !c.IsSupportedDomain(normalizedName) {
		return nil, &DomainNotSupportedError{DomainName: normalizedName}
	}
	namehash := kns.NameHash(normalizedName)
	tokenID := namehash.Big()
	data, err := c.proxyReader.GetData(&bind.CallOpts{Pending: false}, keys, tokenID)
	if err != nil {
		return nil, err
	}
	if data.Owner == unsZeroAddress {
		return nil, &DomainNotRegisteredError{DomainName: normalizedName}
	}
	if data.Resolver == unsZeroAddress {
		return nil, &DomainNotConfiguredError{DomainName: normalizedName}
	}

	return &data, nil
}

func (c *Uns) Records(domainName string, keys []string) (map[string]string, error) {
	data, err := c.Data(domainName, keys)
	if err != nil {
		return nil, err
	}
	allRecords := make(map[string]string, len(keys))
	for index, key := range keys {
		allRecords[key] = data.Values[index]
	}
	return allRecords, nil
}

func (c *Uns) Record(domainName string, key string) (string, error) {
	records, err := c.Records(domainName, []string{key})
	if err != nil {
		return "", err
	}
	return records[key], nil
}

func (c *Uns) Addr(domainName string, ticker string) (string, error) {
	key, err := buildCryptoKey(ticker)
	if err != nil {
		return "", err
	}
	value, err := c.Record(domainName, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *Uns) AddrVersion(domainName string, ticker string, version string) (string, error) {
	key, err := buildCryptoKeyVersion(ticker, version)
	if err != nil {
		return "", err
	}
	value, err := c.Record(domainName, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *Uns) Email(domainName string) (string, error) {
	value, err := c.Record(domainName, emailKey)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (c *Uns) Resolver(domainName string) (string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return "", err
	}

	return data.Resolver.String(), nil
}

func (c *Uns) Owner(domainName string) (string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return "", err
	}

	return data.Owner.String(), nil
}

func (c *Uns) IpfsHash(domainName string) (string, error) {
	records, err := c.Records(domainName, ipfsKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, ipfsKeys), nil
}

func (c *Uns) HTTPUrl(domainName string) (string, error) {
	records, err := c.Records(domainName, redirectUrlKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, redirectUrlKeys), nil
}

func (c *Uns) getAllKeysFromContractEvents(contract *registry.Contract, eventsStartingBlock uint64, domainName string) ([]string, error) {
	var allKeys []string
	normalizedName := normalizeName(domainName)
	namehash := kns.NameHash(normalizedName)
	resetRecordsIterator, err := contract.FilterResetRecords(&bind.FilterOpts{Start: eventsStartingBlock}, []*big.Int{namehash.Big()})
	if err != nil {
		return nil, err
	}
	newKeyEventsStartingBlock := eventsStartingBlock
	for resetRecordsIterator.Next() {
		if resetRecordsIterator.Error() != nil {
			return nil, err
		}
		newKeyEventsStartingBlock = resetRecordsIterator.Event.Raw.BlockNumber
	}
	newKeyIterator, err := contract.FilterNewKey(&bind.FilterOpts{Start: newKeyEventsStartingBlock}, []*big.Int{namehash.Big()}, []string{})
	if err != nil {
		return nil, err
	}
	for newKeyIterator.Next() {
		if newKeyIterator.Error() != nil {
			return nil, err
		}
		allKeys = append(allKeys, newKeyIterator.Event.Key)
	}
	if len(allKeys) == 0 {
		for key := range c.supportedKeys {
			allKeys = append(allKeys, key)
		}
	}
	return allKeys, err
}

func (c *Uns) AllRecords(domainName string) (map[string]string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return nil, err
	}
	var allKeys []string
	if data.Resolver == cnsMainnetDefaultResolver || data.Resolver == unsMainnetProxyReader {
		contract, err := registry.NewContract(data.Resolver, c.contractBackend)
		if err != nil {
			return nil, err
		}
		allKeys, err = c.getAllKeysFromContractEvents(contract, unsEventsStartingBlock, domainName)
		if err != nil {
			return nil, err
		}
	} else {
		for key := range c.supportedKeys {
			allKeys = append(allKeys, key)
		}
	}
	recordsData, err := c.Data(domainName, allKeys)
	if err != nil {
		return nil, err
	}
	allRecords := make(map[string]string)
	for index, key := range allKeys {
		if len(recordsData.Values[index]) > 0 {
			allRecords[key] = recordsData.Values[index]
		}
	}

	return allRecords, nil
}

func (c *Uns) DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error) {
	keys, err := dnsTypesToCryptoRecordKeys(types)
	if err != nil {
		return nil, err
	}
	records, err := c.Records(domainName, keys)
	if err != nil {
		return nil, err
	}
	dnsRecords, err := cryptoRecordsToDNS(records)
	if err != nil {
		return nil, err
	}

	return dnsRecords, nil
}

func (c *Uns) IsSupportedDomain(domainName string) bool {
	return !strings.HasSuffix(domainName, ".zil")
}

func (c *Uns) TokenURI(domainName string) (string, error) {
	normalizedName := normalizeName(domainName)
	if !c.IsSupportedDomain(normalizedName) {
		return "", &DomainNotSupportedError{DomainName: normalizedName}
	}
	namehash := kns.NameHash(normalizedName)
	tokenUri, err := c.tokenUriByNamehash(namehash)
	if err != nil {
		return "", err
	}

	return tokenUri, nil
}

func (c *Uns) TokenURIMetadata(domainName string) (TokenMetadata, error) {
	tokenUri, err := c.TokenURI(domainName)
	if err != nil {
		return TokenMetadata{}, err
	}
	metadata, err := c.tokenMetadataByUri(tokenUri)
	if err != nil {
		return TokenMetadata{}, err
	}
	return metadata, nil
}

func (c *Uns) Unhash(domainHash string) (string, error) {
	namehash := common.HexToHash(domainHash)
	tokenUri, err := c.tokenUriByNamehash(namehash)
	if err != nil {
		return "", err
	}
	metadata, err := c.tokenMetadataByUri(tokenUri)
	if err != nil {
		return "", err
	}
	domainName := normalizeName(metadata.Name)
	expectedNamehash := kns.NameHash(domainName)
	if namehash != expectedNamehash {
		return "", &InvalidDomainNameReturnedError{
			DomainName: domainName,
			Namehash:   domainHash,
		}
	}

	return domainName, nil
}

func (c *Uns) tokenUriByNamehash(namehash common.Hash) (string, error) {
	tokenId := namehash.Big()
	tokenUri, err := c.proxyReader.TokenURI(&bind.CallOpts{Pending: false}, tokenId)
	if err != nil {
		if err.Error() == vm.ErrExecutionReverted.Error() {
			return "", &DomainNotRegisteredError{Namehash: namehash.String()}
		}
		return "", err
	}

	return tokenUri, nil
}

func (c *Uns) tokenMetadataByUri(tokenUri string) (TokenMetadata, error) {
	metadataResponse, err := c.metadataClient.Get(tokenUri)
	if err != nil {
		return TokenMetadata{}, err
	}
	defer metadataResponse.Body.Close()
	var returnedMetadata TokenMetadata
	err = json.NewDecoder(metadataResponse.Body).Decode(&returnedMetadata)
	if err != nil {
		return TokenMetadata{}, err
	}

	return returnedMetadata, nil
}
