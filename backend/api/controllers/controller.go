package controllers

import (
	"fmt"
	"net/http"
	"pm-with-essence/api/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// GETRequest ShowAccount
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400  {object}  errors.StatusError
// @Router       /api/get [get]
func (ct Controller) GETRequest(c *gin.Context) {
	err := ct.service.Run()
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
