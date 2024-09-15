package config

import (
	"fmt"
	"os"
)

const (
	DefaultConfigurationLocation = "./config.yaml"
	CouldNotOpenFile             = "Could not open config file"
)

type (
	DatabaseConfig struct {
		Host         string `yaml:"Host"`
		Port         int    `yaml:"Port"`
		Username     string `yaml:"Username"`
		Password     string `yaml:"Password"`
		DriverName   string `yaml:"DriverName"`
		DatabaseName string `yaml:"DatabaseName"`
	}

	LoggingConfiguration struct {
		HandlerType string `yaml:"HandlerType"`
		WriterType  string `yaml:"WriterType"`
		Level       string `yaml:"Level"`
		FileName    string `yaml:"FileName,omitempty"`
	}

	ServerConfig struct {
		Port int `yaml:"Port"`
	}

	Configuration struct {
		DatabaseConfiguration DatabaseConfig       `yaml:"DatabaseConfiguration"`
		LoggingConfiguration  LoggingConfiguration `yaml:"LoggingConfiguration"`
		ServerConfiguration   ServerConfig         `yaml:"ServerConfiguration"`
		Options               map[string]string    `yaml:"Options"`
	}
)

func InitConfiguration() (configuration *Configuration) {
	// ToDo: This function needs to have more configuration checking.
	file, err := os.Open(DefaultConfigurationLocation)
	if err != nil {
		_ = fmt.Errorf(CouldNotOpenFile)
	}
	configuration = &Configuration{}
	if err := configuration.LoadConfiguration(file); err != nil {
		_ = fmt.Errorf(CouldNotOpenFile)
	}
	return
}
