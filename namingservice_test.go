package resolution

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unstoppabledomains/resolution-go/namingservice"
)

func TestEnforceImplementInterface(t *testing.T) {
	t.Parallel()
	assert.Implements(t, (*NamingService)(nil), &Zns{provider: nil})
	assert.Implements(t, (*NamingService)(nil), &Uns{
		proxyReader:     nil,
		supportedKeys:   nil,
		contractBackend: nil,
	})
}

func TestDetectNamingServiceType(t *testing.T) {
	t.Parallel()
	var serviceType string
	serviceType, err := DetectNamingService("test.zil")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.ZNS, serviceType)

	serviceType, err = DetectNamingService("test.crypto")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.CNS, serviceType)

	serviceType, err = DetectNamingService("test.asdasdas")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.CNS, serviceType)
}

func TestDetectNamingServiceTypeInvalidDomain(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotSupportedError
	_, err := DetectNamingService("aaaazzsd..")
	assert.ErrorAs(t, err, &expectedError)

	_, err = DetectNamingService("aaaazzsd")
	assert.ErrorAs(t, err, &expectedError)
}
