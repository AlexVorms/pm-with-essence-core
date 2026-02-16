package main

import (
	"fmt"
	"log"
	"pm-with-essence/api"
	"pm-with-essence/api/controllers"
	"pm-with-essence/config"
	_ "pm-with-essence/docs"
)

// @title           Project manager with Essence core
// @version         1.0
// @description     Серверная часть прототипа системы управления проектами с ядром Essence
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
func main() {
	//set config
	cfg, err := config.SetConfig()
	if err != nil {
		log.Fatal("Не удалось загрузить конфиг:", err)
	}

	fmt.Printf("Запуск на порту: %d\n", cfg.Port)
	fmt.Printf("БД Хост: %s, Пользователь: %s\n", cfg.Database.Host, cfg.Database.Username)
	fmt.Println(cfg.Database.Password)

	//init controllers
	controller := controllers.NewController()

	////init routes
	handler := api.NewRouter(*controller)
	fmt.Printf("Swagger running on http://localhost:8080/swagger/index.html")
	handler.InitRoutes()

}
