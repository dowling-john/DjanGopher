package config

type (
	DatabaseConfig struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
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
	configuration = &Configuration{}
	configuration.LoadConfiguration()
	return
}
