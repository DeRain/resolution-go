package resolution

import (
	"encoding/json"
	"strings"

	"github.com/unstoppabledomains/resolution-go/namingservice"

	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
)

// Zns is a naming service handles .zil domains resolution.
type Zns struct {
	provider ZnsProvider
}

// ZnsBuilder is a builder to setup and build instance of Zns service.
type ZnsBuilder interface {
	// SetProvider set Zilliqa blockchain provider to communicate with ZNS registry
	SetProvider(provider ZnsProvider) ZnsBuilder
	// Build Zns instance
	Build() (*Zns, error)
}

type znsBuilder struct {
	provider ZnsProvider
}

// ZnsProvider ZnsProvider
type ZnsProvider interface {
	GetSmartContractSubState(contractAddress string, params ...interface{}) (string, error)
}

// ZnsDomainState State of ZNS domain
type ZnsDomainState struct {
	Resolver string
	Owner    string
	Records  map[string]string
}

type registrySubState struct {
	Result map[string]map[string]struct {
		Arguments   []string
		Argtypes    []string
		Constructor string
	}
}

type resolverSubState struct {
	Result map[string]map[string]string
}

const znsDefaultProvider = "https://api.zilliqa.com"
const znsMainnetRegistry = "9611c53BE6d1b32058b2747bdeCECed7e1216793"
const znsContractField = "records"
const znsZeroAddress = "0x0000000000000000000000000000000000000000"

// NewZnsBuilder Creates ZNS builder instance.
func NewZnsBuilder() ZnsBuilder {
	return &znsBuilder{}
}

// SetProvider set Zilliqa blockchain provider to communicate with ZNS registry.
func (zb *znsBuilder) SetProvider(provider ZnsProvider) ZnsBuilder {
	zb.provider = provider
	return zb
}

// Build Zns instance
func (zb *znsBuilder) Build() (*Zns, error) {
	if zb.provider == nil {
		zb.provider = provider.NewProvider(znsDefaultProvider)
	}

	return &Zns{provider: zb.provider}, nil
}

// State Get raw data attached to domain.
func (z *Zns) State(domainName string) (*ZnsDomainState, error) {
	normalizedName := normalizeName(domainName)
	isSupported, err := z.IsSupportedDomain(normalizedName)
	if err != nil {
		return nil, err
	}
	if !isSupported {
		return nil, &DomainNotSupportedError{DomainName: normalizedName}
	}
	namehash, err := ZnsNameHash(domainName)
	if err != nil {
		return nil, err
	}
	response, err := z.provider.GetSmartContractSubState(znsMainnetRegistry, znsContractField, []string{namehash})
	if err != nil {
		return nil, err
	}

	var registryState registrySubState
	err = json.Unmarshal([]byte(response), &registryState)
	if err != nil {
		return nil, err
	}
	registryValues := registryState.Result[znsContractField][namehash].Arguments
	if len(registryValues) == 0 {
		return nil, &DomainNotRegisteredError{DomainName: domainName}
	}
	owner, resolver := registryValues[0], registryValues[1]
	if owner == znsZeroAddress {
		return nil, &DomainNotRegisteredError{DomainName: domainName}
	}
	if resolver == znsZeroAddress {
		return nil, &DomainNotConfiguredError{DomainName: domainName}
	}

	response, err = z.provider.GetSmartContractSubState(strings.TrimPrefix(resolver, "0x"), znsContractField, []string{})
	if err != nil {
		return nil, err
	}
	var resolverState resolverSubState
	err = json.Unmarshal([]byte(response), &resolverState)
	if err != nil {
		return nil, err
	}
	records := resolverState.Result[znsContractField]

	return &ZnsDomainState{Owner: owner, Resolver: resolver, Records: records}, nil
}

func (z *Zns) Records(domainName string, keys []string) (map[string]string, error) {
	state, err := z.State(domainName)
	if err != nil {
		return nil, err
	}
	records := make(map[string]string, len(keys))
	for _, recordKey := range keys {
		records[recordKey] = state.Records[recordKey]
	}

	return records, err
}

func (z *Zns) Record(domainName string, key string) (string, error) {
	records, err := z.Records(domainName, []string{key})
	if err != nil {
		return "", nil
	}
	return records[key], nil
}

func (z *Zns) Owner(domainName string) (string, error) {
	state, err := z.State(domainName)
	if err != nil {
		return "", err
	}

	return state.Owner, err
}

func (z *Zns) Resolver(domainName string) (string, error) {
	state, err := z.State(domainName)
	if err != nil {
		return "", err
	}

	return state.Resolver, err
}

func (z *Zns) Addr(domainName string, ticker string) (string, error) {
	key, err := buildCryptoKey(ticker)
	if err != nil {
		return "", err
	}
	value, err := z.Record(domainName, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (z *Zns) AddrVersion(domainName string, ticker string, version string) (string, error) {
	key, err := buildCryptoKeyVersion(ticker, version)
	if err != nil {
		return "", err
	}
	value, err := z.Record(domainName, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (z *Zns) Email(domainName string) (string, error) {
	value, err := z.Record(domainName, emailKey)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (z *Zns) AllRecords(domainName string) (map[string]string, error) {
	state, err := z.State(domainName)
	if err != nil {
		return nil, err
	}

	return state.Records, err
}

func (z *Zns) IpfsHash(domainName string) (string, error) {
	records, err := z.Records(domainName, ipfsKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, ipfsKeys), nil
}

func (z *Zns) HTTPUrl(domainName string) (string, error) {
	records, err := z.Records(domainName, redirectUrlKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, redirectUrlKeys), nil
}

func (z *Zns) DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error) {
	keys, err := dnsTypesToCryptoRecordKeys(types)
	if err != nil {
		return nil, err
	}
	records, err := z.Records(domainName, keys)
	if err != nil {
		return nil, err
	}
	dnsRecords, err := cryptoRecordsToDNS(records)
	if err != nil {
		return nil, err
	}

	return dnsRecords, nil
}

func (z *Zns) IsSupportedDomain(domainName string) (bool, error) {
	return strings.HasSuffix(domainName, ".zil"), nil
}

func (z *Zns) TokenURI(_ string) (string, error) {
	return "", &MethodIsNotSupportedError{NamingServiceName: namingservice.ZNS}
}

func (z *Zns) TokenURIMetadata(_ string) (TokenMetadata, error) {
	return TokenMetadata{}, &MethodIsNotSupportedError{NamingServiceName: namingservice.ZNS}
}

func (z *Zns) Unhash(_ string) (string, error) {
	return "", &MethodIsNotSupportedError{NamingServiceName: namingservice.ZNS}
}
