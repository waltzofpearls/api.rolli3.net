package main

import "github.com/kelseyhightower/envconfig"

const (
	dev  = "dev"
	prod = "prod"
)

type config struct {
	Env     string `envconfig:"ENV" default:"dev"`
	Listen  string `envconfig:"LISTEN" default:":9443"`
	TLSKey  string `envconfig:"TLS_KEY" default:"assets/local.api.rolli3.net.key.pem"`
	TLSCert string `envconfig:"TLS_CERT" default:"assets/local.api.rolli3.net.cert.pem"`
}

func newConfig() *config {
	return new(config)
}

func (c *config) load() error {
	return envconfig.Process("", c)
}

func (c *config) isEnv(name string) bool {
	return c.Env == name
}
