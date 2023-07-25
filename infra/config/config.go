package config

type Config struct {
	Type                    string `json:"type" env:"type"`
	ProjectId               string `json:"project_id" env:"project_id"`
	PrivateKeyId            string `json:"private_key_id" env:"private_key_id"`
	PrivateKey              string `json:"private_key" env:"private_key"`
	ClientEmail             string `json:"client_email" env:"client_email"`
	ClientId                string `json:"client_id" env:"client_id"`
	AuthUri                 string `json:"auth_uri" env:"auth_uri"`
	TokenUri                string `json:"token_uri" env:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url" env:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url" env:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain" env:"universe_domain"`
}
