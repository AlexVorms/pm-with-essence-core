package routes

import (
	"pm-with-essence/cmd/api/controllers"
	"pm-with-essence/config"
	_ "pm-with-essence/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	userController    controllers.UserController
	taskController    controllers.TaskController
	projectController controllers.ProjectController
	boardController   controllers.BoardController
	columnController  controllers.ColumnController
}

func NewRouter(
	userController controllers.UserController,
	taskController controllers.TaskController,
	projectController controllers.ProjectController,
	boardController controllers.BoardController,
	columnController controllers.ColumnController) *Router {
	return &Router{
		userController:    userController,
		taskController:    taskController,
		projectController: projectController,
		boardController:   boardController,
		columnController:  columnController}
}
func (r *Router) InitRoutes(cfg config.RouterConfig) (*gin.Engine, error) {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type"}

	router.Use(cors.New(config))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	userController := router.Group("/")
	{
		userController.POST("register", r.userController.Register)
		userController.POST("login", r.userController.Login)
		userController.GET("profile", r.userController.Profile)
	}
	taskController := router.Group("/task")
	{
		taskController.POST("/:columnId", r.taskController.CreateTask)
		taskController.GET("", r.taskController.GetAllTasks)
		taskController.GET("/:taskId", r.taskController.GetTask)
		taskController.DELETE("/:taskId", r.taskController.DeleteTask)
		taskController.PUT("/:taskId", r.taskController.UpdateTask)
		taskController.PUT("/column/:columnId/task/:taskId", r.taskController.ChangeTaskColumn)
		taskController.PUT("/:taskId/complete", r.taskController.FinishTask)
	}
	projectController := router.Group("/project")
	{
		projectController.POST("", r.projectController.CreateProject)
		projectController.GET("", r.projectController.GetAllProjects)
		projectController.GET("/:projectId", r.projectController.GetProject)
		projectController.DELETE("/:projectId", r.projectController.DeleteProject)
		projectController.PUT("/:projectId", r.projectController.UpdateProject)
	}
	boardController := router.Group("/board")
	{
		boardController.POST("/:projectId", r.boardController.CreateBoard)
		boardController.GET("/all-boards/:projectId", r.boardController.GetBoards)
		boardController.GET("/:boardId", r.boardController.GetBoard)
		boardController.DELETE("/:boardId", r.boardController.DeleteBoard)
		boardController.PUT("/:boardId", r.boardController.UpdateBoard)
	}
	columnController := router.Group("/column")
	{
		columnController.POST("/:boardID", r.columnController.CreateColumn)
		columnController.PUT("/:columnID", r.columnController.UpdateColumn)
		columnController.DELETE("/:columnID", r.columnController.DeleteColumn)
		columnController.PUT("", r.columnController.UpdateColumnOrder)
	}
	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})
	err := router.Run(cfg.Port)
	if err != nil {
		return nil, err
	}

	return router, nil
}
