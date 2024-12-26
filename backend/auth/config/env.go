package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort string `mapstructure:"SERVER_PORT"`

	// database
	DBHost    string `mapstructure:"DB_HOST"`
	DBName    string `mapstructure:"DB_NAME"`
	DBPort    uint16 `mapstructure:"DB_PORT"`
	DBUser    string `mapstructure:"DB_USER"`
	DBUserPwd string `mapstructure:"DB_USER_PWD"`

	// jWT
	JWTSecret string `mapstructure:"JWT_SECRET"`
}

func NewEnv(filename string) *Env {
	env := Env{}

	viper.SetConfigFile(filename)

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Error reading enviroment file")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Error loading environment file", err)
	}

	return &env
}
