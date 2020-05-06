package config

import (
	"github.com/gargakshit/paperplane-server/model"
)

// GetDefaultConfig creates and returns the default config
func GetDefaultConfig() model.Config {
	defaultConfig := &model.Config{}
	defaultConfig.SetDefaults()

	return *defaultConfig
}
