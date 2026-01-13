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

func LoadConfig() (c Config, err error) {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	c.Port = viper.GetString("PORT")
	c.DBUrl = viper.GetString("DB_URL")

	c.Auth = Auth{
		Name:     viper.GetString("USER_NAME"),
		Password: viper.GetString("PASSWORD"),
	}

	return c, nil
}
