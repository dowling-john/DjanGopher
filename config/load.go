package config

import (
	"fmt"
	"github.com/dowling-john/DjanGopher/errors"
	"gopkg.in/yaml.v2"
	"os"
)

const DefaultConfigurationLocation = "./config.yaml"

type (
	DatabaseConfig struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	ServerConfig struct {
		Port int `yaml:"Port"`
	}

	Configuration struct {
		DatabaseConfiguration DatabaseConfig    `yaml:"DatabaseConfiguration"`
		ServerConfiguration   ServerConfig      `yaml:"ServerConfiguration"`
		Options               map[string]string `yaml:"Options"`
	}
)

// LoadConfiguration
// Looks for a yaml file in the root of the application and loads the content into the Options variable this configuration
// type, this type is then returned to the caller, if the file is not found it will cause the application to close.
func (c *Configuration) LoadConfiguration() {
	fmt.Println("Finding configuration...")
	yamlFile, err := os.ReadFile(DefaultConfigurationLocation)
	errors.LogAnyErrorAndExit(err)
	fmt.Printf("Found configuration in %v...\n", DefaultConfigurationLocation)

	fmt.Println("Loading YAML from configuration file...")
	err = yaml.Unmarshal(yamlFile, c)
	errors.LogAnyErrorAndExit(err)
	fmt.Println("Configuration loaded")
}

func (d *DatabaseConfig) FormatDsn() string {
	return ""
}
