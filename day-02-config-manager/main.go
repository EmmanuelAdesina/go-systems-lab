package main

import (
	"fmt"
	"time"
)

// Config represents the configuration for a web service.
type Config struct {
	Host string
	Port int
	TLSEnabled bool
	CertFile string
	KeyFile string
	Timeout time.Duration
	LogLevel string
}

// ChangePort updates the listening port.
func (c *Config) ChangePort(port int) {
	c.Port = port
}

// EnableTLS enables TLS and sets the certificate and
// key file paths.
func (c *Config) EnableTLS(certFile, keyFile string) {
	c.TLSEnabled = true
	c.CertFile = certFile
	c.KeyFile = keyFile
}

func main() {
	// Create a default config
	cfg := Config{
		Host: "0.0.0.0",
		Port: 8080,
		TLSEnabled: false,
		CertFile: "",
		KeyFile: "",
		Timeout: 30 * time.Second,
		LogLevel: "info",
	}

	fmt.Println("=== Before changes ===")
	fmt.Printf("%+v\n", cfg)

	// Modify the config
	cfg.ChangePort(8443) 
	cfg.EnableTLS("/etc/certs/server.crt", "/etc/certs/server.key")
	fmt.Println("\n=== After changes ===")
	fmt.Printf("%+v\n", cfg)
}
