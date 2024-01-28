package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	ConnStr                string `mapstructure:"CONN_STR"`
	RabbitDialStr          string `mapstructure:"RABBIT_DIAL_STR"`
	RabbitExchange         string `mapstructure:"RABBIT_EXCHANGE"`
	RabbitKey              string `mapstructure:"RABBIT_KEY"`
	MercadoPagoAccessToken string `mapstructure:"MP_ACCESS_TOKEN"`
	WebhookNotification    string `mapstructure:"MP_WEBHOOK_NOTIFICATION"`
}

func LoadConfig(path string) (*Config, error) {
	if os.Getenv("ENVIRONMENT") == "PROD" {
		cfg := &Config{
			ConnStr:                os.Getenv("CONN_STR"),
			RabbitDialStr:          os.Getenv("RABBIT_DIAL_STR"),
			RabbitExchange:         os.Getenv("RABBIT_EXCHANGE"),
			RabbitKey:              os.Getenv("RABBIT_KEY"),
			MercadoPagoAccessToken: os.Getenv("MP_ACCESS_TOKEN"),
			WebhookNotification:    os.Getenv("MP_WEBHOOK_NOTIFICATION"),
		}
		return cfg, nil
	}
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
