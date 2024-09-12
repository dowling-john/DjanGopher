package config

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
		WriterType string `yaml:"WriterType"`
		Level      string `yaml:"Level"`
		FileName   string `yaml:"FileName,omitempty"`
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

	configuration = &Configuration{}
	configuration.LoadConfiguration()
	return
}
