package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
	Auth  Auth
}

type Auth struct {
	Name     string `mapstructure:"USER_NAME"`
	Password string `mapstructure:"PASSWORD"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	config.Port = viper.GetString("PORT")
	config.DBUrl = viper.GetString("DB_URL")

	config.Auth = Auth{
		Name:     viper.GetString("USER_NAME"),
		Password: viper.GetString("PASSWORD"),
	}

	return config, nil
}
