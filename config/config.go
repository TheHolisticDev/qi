package config

import "os"

// AppConfig represents all externally configurations for the app
type AppConfig struct {
	Port string
}

// NewConfig returns a new AppConfig with parameters read from the environment or default values if not fount in env
func NewConfig() AppConfig {
	appConfig := AppConfig{
		Port: "1337",
	}

	if port, isSet := os.LookupEnv("PORT"); isSet {
		appConfig.Port = port
	}

	return appConfig
}
