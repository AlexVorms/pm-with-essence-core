package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Router    RouterConfig `mapstructure:"router"`
	Database  Database     `mapstructure:"db"`
	JWTSecret string       `mapstructure:"jwt_secret"`
}
type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}
type RouterConfig struct {
	Port string `mapstructure:"port"`
}

func SetConfig() (config Config, err error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()
	viper.BindEnv("jwt_secret", "JWT_SECRET")
	viper.BindEnv("db.password", "DB_PASSWORD")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Println("Конфигурационный файл не найден, используем окружение")
		}
	}

	err = viper.Unmarshal(&config)
	return config, err
}
