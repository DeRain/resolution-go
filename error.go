package resolution

// DomainNotRegisteredError Error when domain is missing an owner
type DomainNotRegisteredError struct {
	DomainName string
	Namehash   string
}

// DomainNotConfiguredError Error when domain does not have a resolver set
type DomainNotConfiguredError struct {
	DomainName string
}

// DomainNotSupportedError Error when domain is not supported by the naming service
type DomainNotSupportedError struct {
	DomainName string
}

// MethodIsNotSupportedError Error when naming services does not support called method
type MethodIsNotSupportedError struct {
	NamingServiceName string
}

// InvalidDomainNameReturnedError Error when ERC721 metadata provides returns incorrect domain name
type InvalidDomainNameReturnedError struct {
	Namehash   string
	DomainName string
}

func (e *DomainNotRegisteredError) Error() string {
	return "Domain is not registered. Domain name: " + e.DomainName + ". Namehash: " + e.Namehash
}
func (e *DomainNotConfiguredError) Error() string {
	return e.DomainName + " does not have configured Resolver"
}
func (e *DomainNotSupportedError) Error() string {
	return e.DomainName + " is not supported by naming service"
}

func (e *MethodIsNotSupportedError) Error() string {
	return "Method is not supported in " + e.NamingServiceName + " naming service"
}

func (e *InvalidDomainNameReturnedError) Error() string {
	return "Domain name " + e.DomainName + " was returned from metadata provider which namehash does not match with requested namehash: " + e.Namehash
}
