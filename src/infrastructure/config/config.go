package config

import (
	"fmt"

	logger "gingoskeleton/src/infrastructure/providers"
)

type Config struct {
	// Application
	AppName    string `env:"APP_NAME" envDefault:"Gingoskeleton"`
	AppEnv     string `env:"APP_ENV" envDefault:"development"`
	AppPort    string `env:"APP_PORT" envDefault:"8080"`
	AppVersion string `env:"APP_VERSION" envDefault:"0.0.1"`
	EnableLog  bool   `env:"ENABLE_LOG" envDefault:"true"`
	DebugLog   bool   `env:"DEBUG_LOG" envDefault:"true"`
}

func (c *Config) ToMap() map[string]string {
	return map[string]string{
		"APP_NAME":    c.AppName,
		"APP_ENV":     c.AppEnv,
		"APP_PORT":    c.AppPort,
		"APP_VERSION": c.AppVersion,
		"ENABLE_LOG":  fmt.Sprintf("%t", c.EnableLog),
		"DEBUG_LOG":   fmt.Sprintf("%t", c.DebugLog),
	}
}

func NewConfig() *Config {
	config, err := LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration")
		logger.Logger.Panic("Error loading configuration", err)
	}

	logger.Logger.Setup(config.EnableLog, config.DebugLog)
	return config
}

var ConfigInstance *Config

func init() {
	ConfigInstance = NewConfig()
}
