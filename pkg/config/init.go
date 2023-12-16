package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig(absoluteWorkSpacePath string) {
	viper.SetConfigName("config")              // name of config file (without extension)
	viper.SetConfigType("yaml")                // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(absoluteWorkSpacePath) // path to look for the config file in
	err := viper.ReadInConfig()                // Find and read the config file
	if err != nil {                            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
