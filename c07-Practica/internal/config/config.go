package config

import "github.com/spf13/viper"

type Config struct {
	ServerAddress string
	DatabaseURL   string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		ServerAddress: viper.GetString("server_address"),
		DatabaseURL:   viper.GetString("database_url"),
	}

	return cfg, nil

}
