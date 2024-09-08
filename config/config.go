package config

func InitConfiguration() (configuration *Configuration) {
	configuration = &Configuration{}
	configuration.LoadConfiguration()
	return
}
