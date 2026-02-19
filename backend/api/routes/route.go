package routes

import (
	"pm-with-essence/api/controllers"
	"pm-with-essence/config"
	_ "pm-with-essence/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	controller controllers.Controller
}

func NewRouter(controller controllers.Controller) *Router {
	return &Router{controller: controller}
}
func (r *Router) InitRoutes(cfg config.RouterConfig) (*gin.Engine, error) {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type"}

	router.Use(cors.New(config))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	controller := router.Group("/api")
	{
		controller.GET("/get", r.controller.GETRequest)
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
