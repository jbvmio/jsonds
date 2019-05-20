package jsonds

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds the configuration.
type Config struct {
	Name        string
	LogLevel    string
	HTTPAddress string
	Endpoints   map[string]string
}

// GetConfig reads in the config file.
func GetConfig(filePath string) *Config {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Unable to Read Config: %v\n", err)
	}
	viper.SetDefault(`http.address`, `:8080`)
	viper.SetDefault(`loglevel`, `info`)
	return &Config{
		LogLevel:    viper.GetString(`loglevel`),
		HTTPAddress: viper.GetString(`http.address`),
		Endpoints:   viper.GetStringMapString(`endpoints`),
	}
}
