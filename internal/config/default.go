package config

// Default returns the default configuration
func Default() Config {
	return Config{
		ExportInterval: 10,
		ServicePort:    "8080",
	}
}
