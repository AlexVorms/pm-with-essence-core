package main

import (
	"fmt"
	"log"
	"pm-with-essence/cmd/api/controllers"
	"pm-with-essence/cmd/api/routes"
	"pm-with-essence/cmd/repository"
	"pm-with-essence/cmd/service/authService"
	"pm-with-essence/cmd/service/boardService"
	"pm-with-essence/cmd/service/columnService"
	"pm-with-essence/cmd/service/projectService"
	"pm-with-essence/cmd/service/taskService"
	userService "pm-with-essence/cmd/service/userService"
	"pm-with-essence/config"
	_ "pm-with-essence/docs"
	"pm-with-essence/internal/database"

	"github.com/joho/godotenv"
)

// @title           Project manager with Essence core
// @version         1.0
// @description     Серверная часть прототипа системы управления проектами с ядром Essence

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Введите JWT токен в формате: Bearer {token}

func main() {
	//set config
	_ = godotenv.Load()

	cfg, err := config.SetConfig()
	if err != nil {
		log.Fatal("Не удалось загрузить конфиг:", err)
	}

	fmt.Printf("Запуск на порту: %d\n", cfg.Router.Port)
	fmt.Printf("БД Хост: %s, Пользователь: %s\n", cfg.Database.Host, cfg.Database.Username)

	//init database
	db, err := database.ConnectDb(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
	//}

	//init repositories
	projectStorage := repository.NewProjectStorage(db)
	userStorage := repository.NewUserStorage(db)
	taskStorage := repository.NewTaskStorage(db)
	boardStorage := repository.NewBoardStorage(db)
	columnStorage := repository.NewColumnStorage(db)

	//init services
	newProjectService := projectService.NewProjectService(projectStorage)
	newUserService := userService.NewUserService(userStorage)
	newJWTService := authService.NewJwtService(cfg.JWTSecret)
	newAuthService := authService.NewAuthService(userStorage, newJWTService)
	newTaskService := taskService.NewTaskService(taskStorage, columnStorage)
	newColumnService := columnService.NewColumnService(columnStorage, boardStorage)
	newBoardService := boardService.NewBoardService(boardStorage, newColumnService, newProjectService)

	//init controllers
	newProjectController := controllers.NewProjectController(newProjectService)
	userController := controllers.NewUserController(newUserService, newAuthService)
	taskController := controllers.NewTaskController(newTaskService)
	boardController := controllers.NewBoardController(newBoardService)
	columnController := controllers.NewColumnController(newColumnService)

	//init routes
	handler := routes.NewRouter(
		*userController,
		*taskController,
		*newProjectController,
		*boardController,
		*columnController,
		*newJWTService)
	fmt.Printf("Swagger running on http://localhost:8080/swagger/index.html")
	_, err = handler.InitRoutes(cfg.Router)
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
