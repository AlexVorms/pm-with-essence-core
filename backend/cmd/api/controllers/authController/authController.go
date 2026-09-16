package authController

import (
	"fmt"
	"log"
	"net/http"
	"pm-with-essence/cmd/api/model/responseDTO"
	"pm-with-essence/cmd/api/model/userDTO"
	"pm-with-essence/internal/api/app_errors"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService AuthService
}

type AuthService interface {
	Login(model userDTO.LoginDTO) (*string, *app_errors.HttpError)
}

func NewAuthController(authService AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Login
// @Summary login
// @Tags User
// @Description user login
// @Produce json
// @Param loginDTO body userDTO.LoginDTO true "LoginDTO"
// @Success 200 {object} responseDTO.LoginResponse
// @Router       /login [post]
func (u *AuthController) Login(c *gin.Context) {
	var loginDTO userDTO.LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, responseDTO.ErrorResponse{
			Message: "Invalid request body",
		})
		return
	}
	token, err := u.authService.Login(loginDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, responseDTO.ErrorResponse{
			Message: err.Message,
		})
		return
	}
	c.JSON(http.StatusOK, responseDTO.LoginResponse{
		Message:     "Success",
		TokenType:   "Bearer",
		AccessToken: *token,
	})
}

//TODO:Обновление токена
//TODO:Логаут
