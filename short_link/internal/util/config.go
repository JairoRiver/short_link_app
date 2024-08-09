package util

import "github.com/spf13/viper"

// Config stores all configuration of the application
// The values are read by viper from a config file or environment variable.
type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DomainName    string `mapstructure:"DOMAIN_NAME"`
	Environment   string `mapstructure:"ENVIRONMENT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path, name string) (Config, error) {
	config := Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}
