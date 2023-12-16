package config_reader

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig(absoluteWorkSpacePath string) {
	Config = viper.New()
	Config.SetConfigName("config")                            // name of config file (without extension)
	Config.SetConfigType("yaml")                              // REQUIRED if the config file does not have the extension in the name
	Config.AddConfigPath(absoluteWorkSpacePath)               // path to look for the config file in
	Config.AddConfigPath(absoluteWorkSpacePath + "/configs/") // path to look for the config file in

	err := Config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

// GetConfig 返回配置对象
func GetViperConfig() *viper.Viper {
	return Config
}

func GetConfigByKey(key string) string {
	return Config.GetString(key)
}
