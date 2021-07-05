package resolution

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
	"github.com/unstoppabledomains/resolution-go/uns/contracts/registry"
)

type MockedMetadataClient struct {
	Response *http.Response
	Err      error
}

func (m *MockedMetadataClient) SetResponse(resp *http.Response) *MockedMetadataClient {
	m.Response = resp
	return m
}

func (m *MockedMetadataClient) SetError(err error) *MockedMetadataClient {
	m.Err = err
	return m
}

func (m *MockedMetadataClient) Get(_ string) (resp *http.Response, err error) {
	return m.Response, m.Err
}

var uns, _ = NewUnsBuilder().Build()

func TestUnsBuilder(t *testing.T) {
	t.Parallel()
	builder := NewUnsBuilder()
	_, err := builder.Build()
	assert.Nil(t, err)
	assert.NotNil(t, uns.contractBackend)
	assert.NotNil(t, uns.metadataClient)
	assert.NotNil(t, uns.supportedKeys)
	assert.NotNil(t, uns.proxyReader)
}

func TestUnsBuilderSetBackend(t *testing.T) {
	t.Parallel()
	backend, _ := ethclient.Dial("https://rinkeby.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e")
	builder := NewUnsBuilder()
	builder.SetContractBackend(backend)
	uns, err := builder.Build()
	assert.Nil(t, err)
	assert.Equal(t, backend, uns.contractBackend)
}

func TestUnsBuilderSetMetadataClient(t *testing.T) {
	t.Parallel()
	client := &http.Client{}
	builder := NewUnsBuilder()
	builder.SetMetadataClient(client)
	uns, err := builder.Build()
	assert.Nil(t, err)
	assert.Equal(t, client, uns.metadataClient)
}

func TestNewUnsWithSupportedKeys(t *testing.T) {
	t.Parallel()
	unsService, _ := NewUnsBuilder().Build()
	deprecatedKeyName := unsService.supportedKeys["crypto.ETH.address"]
	assert.Equal(t, "ETH", deprecatedKeyName.DeprecatedKeyName)
}

func TestUnsDataValue(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2"
	data, err := uns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
}

func TestUnsData(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2"
	expectedOwner := common.HexToAddress("0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2")
	expectedResolver := common.HexToAddress("0x95AE1515367aa64C462c71e87157771165B1287A")
	data, err := uns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
	assert.Equal(t, expectedOwner, data.Owner)
	assert.Equal(t, expectedResolver, data.Resolver)
}

func TestUnsEmptyDataValues(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	data, _ := uns.Data(testDomain, []string{"empty record"})
	assert.Equal(t, data.Values[0], "")
	assert.Len(t, data.Values, 1)
}

func TestUnsDomainNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	testDomain := "not-registered-long-domain-name.crypto"
	_, err := uns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsDomainNotConfigured(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotConfiguredError
	testDomain := "udtestdev-d0137c.crypto"
	_, err := uns.Data(testDomain, []string{"crypto.XLM.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsRecords(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecords := map[string]string{"crypto.ETH.address": "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2", "crypto.BTC.address": ""}
	records, err := uns.Records(testDomain, []string{"crypto.ETH.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.Equal(t, records, expectedRecords)
}

func TestUnsEmptyRecords(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecords := map[string]string{"crypto.BTC.address": "", "crypto.ETH.address": "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2", "record-not-exist": ""}
	records, err := uns.Records(testDomain, []string{"record-not-exist", "crypto.ETH.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.Equal(t, records, expectedRecords)
}

func TestUnsRecord(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecord := "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2"
	record, err := uns.Record(testDomain, "crypto.ETH.address")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsEmptyRecord(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	record, err := uns.Record(testDomain, "record-not-exist")
	assert.Nil(t, err)
	assert.Empty(t, record)
}

func TestUnsAddr(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2"
	record, err := uns.Addr(testDomain, "ETH")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsAddrLowerCaseTicker(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecord := "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2"
	record, err := uns.Addr(testDomain, "eth")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsEmail(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "testing@example.com"
	record, err := uns.Email(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsResolver(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "0x95AE1515367aa64C462c71e87157771165B1287A"
	record, err := uns.Resolver(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsOwner(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "0x95AE1515367aa64C462c71e87157771165B1287A"
	record, err := uns.Resolver(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsAddrVersion(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2"
	record, err := uns.AddrVersion(testDomain, "USDT", "ERC20")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsIpfs(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "QmRi3PBpUGFnYrCKUoWhntRLfA9PeRhepfFu4Lz21mGd3X"
	record, err := uns.IpfsHash(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsIpfsLegacy(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "QmRi3PBpUGFnYrCKUoWhntRLfA9PeRhepfFu4Lz21mGd3X"
	record, err := uns.IpfsHash(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsHTTPUrl(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-redirect.crypto"
	expectedRecord := "https://example.com/home.html"
	record, err := uns.HTTPUrl(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsHttpUrlLegacy(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-redirect-legacy.crypto"
	expectedRecord := "https://legacy-example.com/home.html"
	record, err := uns.HTTPUrl(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestDotCryptoAllRecords(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecords := map[string]string{"crypto.ETH.address": "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2", "crypto.USDT.version.EOS.address": "karaarishmen", "crypto.USDT.version.ERC20.address": "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2", "crypto.USDT.version.OMNI.address": "1KvzMF2Vjy14d6JGY7dG7vjT5kfpmzSQXM", "crypto.USDT.version.TRON.address": "TRMJfXXbmwb3WFSRKbeRgKsYoD8o1a9xxV", "dns.A": "[\"10.0.0.1\", \"10.0.0.3\"]", "dns.A.ttl": "98", "dns.AAAA": "[]", "ipfs.html.value": "QmRi3PBpUGFnYrCKUoWhntRLfA9PeRhepfFu4Lz21mGd3X", "whois.email.value": "testing@example.com"}
	allRecords, err := uns.AllRecords(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, allRecords)
}

func TestUnsGetAllKeysFromContractEvents(t *testing.T) {
	t.Parallel()

	expectedRecords := []string([]string{"crypto.ETH.address", "crypto.BTC.address"})
	registryContract, err := registry.NewContract(common.HexToAddress("0x7fb83000B8eD59D3eAD22f0D584Df3a85fBC0086"), uns.contractBackend)
	assert.Nil(t, err)
	allKeys, err := uns.getAllKeysFromContractEvents(registryContract, 8775208, "udtestdev-my-new-tls.wallet")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, allKeys)
}

func TestUnsAllRecords(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-my-new-tls.wallet"
	expectedRecords := map[string]string{"crypto.BTC.address": "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", "crypto.ETH.address": "0x6EC0DEeD30605Bcd19342f3c30201DB263291589"}
	allRecords, err := uns.AllRecords(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, allRecords)
}

func TestCnsAllRecordsStandardKeys(t *testing.T) {
	t.Parallel()
	testDomain := "monmouthcounty.crypto"
	expectedRecords := map[string]string{
		"crypto.BTC.address":         "3NwuV8nVT2VKbtCs8evChdiW6kHTHcVpdn",
		"crypto.ETH.address":         "0x1C42088b82f6Fa5fB883A14240C4E066dDFf1517",
		"crypto.LTC.address":         "MTnTNwKikiMi97Teq8XQRabL9SZ4HjnKNB",
		"crypto.ADA.address":         "DdzFFzCqrhsfc3MQvjsLr9BHkaFYeE7BotyTATdETRoSPj6QPiotK4xpcFZk66KVmtr87tvUFTcbTHZRkcdbMR5Ss6jCfzCVtFRMB7WE",
		"ipfs.html.value":            "QmYqX8D8SkaF5YcpaWMyi5xM43UEteFiSNKYsjLcdvCWud",
		"ipfs.redirect_domain.value": "https://abbfe6z95qov3d40hf6j30g7auo7afhp.mypinata.cloud/ipfs/QmYqX8D8SkaF5YcpaWMyi5xM43UEteFiSNKYsjLcdvCWud",
	}
	allRecords, err := uns.AllRecords(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, allRecords)
}

func TestUnsAllRecordsStandardKeys(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-my-new-tls.wallet"

	expectedRecords := map[string]string{"crypto.BTC.address": "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", "crypto.ETH.address": "0x6EC0DEeD30605Bcd19342f3c30201DB263291589"}
	allRecords, err := uns.AllRecords(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, allRecords)
}

func TestUnsDnsA(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-dns.crypto"
	expectedRecords := []dnsrecords.Record{
		{Type: "A", TTL: 1800, Value: "10.0.0.1"},
		{Type: "A", TTL: 1800, Value: "10.0.0.2"},
	}
	dnsRecords, err := uns.DNS(testDomain, []dnsrecords.Type{"A"})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, dnsRecords)
}

func TestUnsDnsCname(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-dns-cname.crypto"
	expectedRecords := []dnsrecords.Record{
		{Type: "CNAME", TTL: 1111, Value: "example.com."},
	}
	dnsRecords, err := uns.DNS(testDomain, []dnsrecords.Type{"CNAME"})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, dnsRecords)
}

func TestUnsDnsGlobalTtl(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-dns-global-ttl.crypto"
	expectedRecords := []dnsrecords.Record{
		{Type: "A", TTL: 1000, Value: "10.0.0.1"},
		{Type: "A", TTL: 1000, Value: "10.0.0.2"},
	}
	dnsRecords, err := uns.DNS(testDomain, []dnsrecords.Type{"A"})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, dnsRecords)
}

func TestUnsIsSupportedDomain(t *testing.T) {
	t.Parallel()

	isSupportedDomain := func(domain string) bool {
		isSupported, _ := uns.IsSupportedDomain(domain)
		return isSupported
	}

	assert.True(t, isSupportedDomain("valid.crypto"))
	assert.True(t, isSupportedDomain("valid.qwdqwd.crypto"))
	assert.False(t, isSupportedDomain("invalid.zil"))
	assert.True(t, isSupportedDomain("invalid.wallet"))
	assert.True(t, isSupportedDomain("invalid.bitcoin"))
	assert.True(t, isSupportedDomain("invalid.x"))
	assert.True(t, isSupportedDomain("invalid.888"))
	assert.True(t, isSupportedDomain("invalid.blockchain"))
	assert.True(t, isSupportedDomain("invalid.dao"))
	assert.True(t, isSupportedDomain("invalid.nft"))
	assert.True(t, isSupportedDomain("invalid.coin"))
	assert.False(t, isSupportedDomain("invalid.com"))
	assert.False(t, isSupportedDomain("radomin-domain.com"))
	assert.False(t, isSupportedDomain("some-domain.net"))
	assert.False(t, isSupportedDomain("some-domain.wiowejfo.qwefwef"))
	assert.False(t, isSupportedDomain("some-domain.wiowejfo.qwd"))
	assert.False(t, isSupportedDomain("some-domain.wiowejfo.zil"))
}

func TestUnsDomainNotRegisteredError(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	_, err := uns.Data("invalid.zil", []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsTokenURI(t *testing.T) {
	t.Parallel()
	tokenURI, err := uns.TokenURI("udtestdev-test.crypto")
	expectedTokenURI := "https://staging-dot-dot-crypto-metadata.appspot.com/metadata/udtestdev-test.crypto"
	assert.Nil(t, err)
	assert.Equal(t, expectedTokenURI, tokenURI)
}

func TestUnsTokenURIDomainIsNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	_, err := uns.TokenURI("unregistered-domain-name.crypto")
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsTokenURIZilDomainIsNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	_, err := uns.TokenURI("invalid.zil")
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsTokenURIMetadata(t *testing.T) {
	t.Parallel()
	expectedMetadata := TokenMetadata{
		Name:        "udtestdev-test.crypto",
		Description: "A .crypto blockchain domain. Use it to resolve your cryptocurrency addresses and decentralized websites.",
		ExternalUrl: "https://staging:Staging-4-Unstoppable-3-2-1@staging.unstoppabledomains.com/search?searchTerm=udtestdev-test.crypto",
		Image:       "https://storage.googleapis.com/dot-crypto-metadata-api/unstoppabledomains_crypto.png",
		Attributes: []TokenMetadataAttribute{
			{
				TraitType: "domain",
				Value:     "udtestdev-test.crypto",
			},
		},
	}
	metadata, err := uns.TokenURIMetadata("udtestdev-test.crypto")
	assert.Nil(t, err)
	assert.Equal(t, expectedMetadata.Name, metadata.Name)
	assert.Equal(t, expectedMetadata.Description, metadata.Description)
	assert.Equal(t, expectedMetadata.ExternalUrl, metadata.ExternalUrl)
	assert.Equal(t, expectedMetadata.Image, metadata.Image)
	assert.Contains(t, metadata.Attributes, expectedMetadata.Attributes[0])
}

func TestUnsTokenURIMetadataNotSupportedDomain(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	_, err := uns.TokenURIMetadata("unregistered-domain-name.crypto")
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsUnhash(t *testing.T) {
	t.Parallel()
	expectedDomainName := "testing.crypto"
	domainName, err := uns.Unhash("0xd52e0f8bfe7e039fddb362c7e00f3628e2dca805f191d8bef74a07ca0e848245")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestUnsUnhashWithout0xPrefix(t *testing.T) {
	t.Parallel()
	expectedDomainName := "testing.crypto"
	domainName, err := uns.Unhash("d52e0f8bfe7e039fddb362c7e00f3628e2dca805f191d8bef74a07ca0e848245")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestUnsUnhashInvalidDomain(t *testing.T) {
	t.Parallel()
	var expectedError *InvalidDomainNameReturnedError
	body, _ := json.Marshal(TokenMetadata{
		Name:            "testing.crypto",
		Description:     "",
		Image:           "",
		ExternalUrl:     "",
		ExternalLink:    "",
		ImageData:       "",
		BackgroundColor: "",
		AnimationUrl:    "",
		YoutubeUrl:      "",
		Attributes:      nil,
	})
	var mockedClient MockedMetadataClient
	mockedClient.SetResponse(&http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(body)),
	})
	mockedClient.SetError(nil)
	unsWithMockedMetadataClient, _ := NewUnsBuilder().SetMetadataClient(&mockedClient).Build()
	domainName, err := unsWithMockedMetadataClient.Unhash("756e4e998dbffd803c21d23b06cd855cdc7a4b57706c95964a37e24b47c10fc9")
	assert.Empty(t, domainName)
	assert.ErrorAs(t, err, &expectedError)
}
