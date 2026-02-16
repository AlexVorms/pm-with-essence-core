package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api [get]
func (ct Controller) GETRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
