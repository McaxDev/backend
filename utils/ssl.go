package utils

import "os"

type SSLConfig struct {
	Enable bool
	Cert   string
	Key    string
}

func LoadSSLConfig(sslc *SSLConfig) {
	if os.Getenv("SSL_ENABLE") == "true" {
		sslc.Enable = true
	}
	sslc.Cert = os.Getenv("SSL_CERT")
	sslc.Key = os.Getenv("SSL_KEY")
}
