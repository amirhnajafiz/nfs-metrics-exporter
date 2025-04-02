package config

// Default returns the default configuration
func Default() Config {
	return Config{
		ServicePort: "8080",
	}
}
