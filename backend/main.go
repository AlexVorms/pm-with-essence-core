package main

import (
	"fmt"
	"log"
	"pm-with-essence/cmd/api/controllers"
	"pm-with-essence/cmd/api/routes"
	"pm-with-essence/cmd/repository"
	"pm-with-essence/cmd/service"
	userService "pm-with-essence/cmd/service/userService"
	"pm-with-essence/config"
	_ "pm-with-essence/docs"
	"pm-with-essence/internal/database"
)

// @title           Project manager with Essence core
// @version         1.0
// @description     Серверная часть прототипа системы управления проектами с ядром Essence
// @securityDefinitions.basic  BasicAuth
func main() {
	//set config
	cfg, err := config.SetConfig()
	if err != nil {
		log.Fatal("Не удалось загрузить конфиг:", err)
	}

	fmt.Printf("Запуск на порту: %d\n", cfg.Router.Port)
	fmt.Printf("БД Хост: %s, Пользователь: %s\n", cfg.Database.Host, cfg.Database.Username)
	fmt.Println(cfg.Database.Password, cfg.Database.DBName)

	//init database
	db, err := database.ConnectDb(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
	//}

	//init repositories
	userStorage := repository.NewUserStorage(db)

	//init services
	serviceR := service.NewService(db)
	newUserService := userService.NewUserService(userStorage)

	//init controllers
	controller := controllers.NewController(serviceR)
	userController := controllers.NewUserController(newUserService)

	//init routes
	handler := routes.NewRouter(*controller,
		*userController)
	fmt.Printf("Swagger running on http://localhost:8080/swagger/index.html")
	_, err = handler.InitRoutes(cfg.Router)
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
