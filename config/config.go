package config

import (
	"github.com/gargakshit/paperplane-server/model"
)

var (
	// GlobalConfig is the global reference to the config
	GlobalConfig *model.Config
)

// GetDefaultConfig creates and returns the default config
func GetDefaultConfig() model.Config {
	defaultConfig := &model.Config{}
	defaultConfig.SetDefaults()

	return *defaultConfig
}
