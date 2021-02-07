package resolution

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

var cns, _ = NewCnsWithDefaultBackend()

func TestNewCnsWithDefaultProvider(t *testing.T) {
	t.Parallel()
	_, err := NewCnsWithDefaultBackend()
	assert.Nil(t, err)
}

func TestNewCns(t *testing.T) {
	t.Parallel()
	backend, _ := ethclient.Dial(DefaultProvider)
	_, err := NewCns(backend)
	assert.Nil(t, err)
}

func TestCnsData(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	data, err := cns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
}

func TestCnsEmptyDataValues(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	data, _ := cns.Data(testDomain, []string{"empty record"})
	assert.Equal(t, data.Values[0], "")
	assert.Len(t, data.Values, 1)
}

func TestCnsDomainNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegistered
	testDomain := "not-registered-long-domain-name.crypto"
	_, err := cns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestCnsDomainNotConfigured(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotConfigured
	testDomain := "reseller-test-paul2.crypto"
	_, err := cns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestCnsRecords(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecords := []string{"0x8aaD44321A86b170879d7A244c1e8d360c99DdA8", "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y"}
	records, err := cns.Records(testDomain, []string{"crypto.ETH.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.ObjectsAreEqual(records, expectedRecords)
}

func TestCnsRecord(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	record, err := cns.Record(testDomain, "crypto.ETH.address")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsAddr(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	record, err := cns.Addr(testDomain, "ETH")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsAddrLowerCaseTicker(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	record, err := cns.Addr(testDomain, "eth")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestEmail(t *testing.T) {
	t.Parallel()
	testDomain := "reseller-test-paul019.crypto"
	expectedRecord := "paul@unstoppabledomains.com"
	record, err := cns.Email(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestResolver(t *testing.T) {
	t.Parallel()
	testDomain := "reseller-test-mago017.crypto"
	expectedRecord := common.HexToAddress("0x878bC2f3f717766ab69C0A5f9A6144931E61AEd3")
	record, err := cns.Resolver(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestAddrVersion(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-usdt.crypto"
	expectedRecord := "0xe7474D07fD2FA286e7e0aa23cd107F8379085037"
	record, err := cns.AddrVersion(testDomain, "USDT", "ERC20")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}
