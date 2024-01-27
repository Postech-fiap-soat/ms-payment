package config

import "github.com/spf13/viper"

type Config struct {
	ConnStr        string `mapstructure:"CONN_STR"`
	RabbitDialStr  string `mapstructure:"RABBIT_DIAL_STR"`
	RabbitExchange string `mapstructure:"RABBIT_EXCHANGE"`
	RabbitKey      string `mapstructure:"RABBIT_KEY"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cfg *Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
