package main

import (
	"fmt"
	"log"
	"pm-with-essence/config"
)

func main() {
	cfg, err := config.SetConfig()
	if err != nil {
		log.Fatal("Не удалось загрузить конфиг:", err)
	}
	fmt.Printf("Запуск на порту: %d\n", cfg.Port)
	fmt.Printf("БД Хост: %s, Пользователь: %s\n", cfg.Database.Host, cfg.Database.Username)
	fmt.Println(cfg.Database.Password)
}
