package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port     int      `mapstructure:"port"`
	Database Database `mapstructure:"db"`
}
type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func SetConfig() (config Config, err error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Конфигурационный файл не найден, используем окружение")
		} else {
			return config, fmt.Errorf("ошибка парсинга файла: %w", err)
		}
	}

	err = viper.Unmarshal(&config)
	return config, err
}
