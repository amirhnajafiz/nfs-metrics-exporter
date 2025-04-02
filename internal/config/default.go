package config

// Default returns the default configuration
func Default() Config {
	return Config{
		DebugMode:      false,
		ExportInterval: 10,
		ServicePort:    "8080",
		SecretKey:      "",
	}
}
