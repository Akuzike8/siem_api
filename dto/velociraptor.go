package dto

type VelociraptorHostsBody struct{
	Hosts string
}

type VelociraptorConfig struct {
	ClientCert      string
	ClientPrivateKey string
	CaCertificate   string
	ApiConnectionString string
}