package delivery_grpc

type Config struct {
	Port               string `yaml:"port"`
	TLSCredentialsPath string `yaml:"tls_credentials_path"`
}
