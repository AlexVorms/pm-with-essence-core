package userController

import (
	"fmt"
	"log"
	"net/http"
	"pm-with-essence/cmd/api/model/responseDTO"
	"pm-with-essence/cmd/api/model/userDTO"
	"pm-with-essence/internal/api/app_errors"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService UserService
}

type UserService interface {
	CreateNewUser(dto userDTO.RegisterDTO) *app_errors.HttpError
}

func NewUserController(userService UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register
// @Summary register
// @Tags User
// @Description user register
// @Produce json
// @Param registerDTO body userDTO.RegisterDTO true "RegisterDTO"
// @Success 200
// @Router       /register [post]
func (u *UserController) Register(c *gin.Context) {

	var registerDTO userDTO.RegisterDTO
	if err := c.ShouldBindJSON(&registerDTO); err != nil {
		log.Printf("error parsing json: %v", err)

		c.JSON(http.StatusBadRequest, responseDTO.ErrorResponse{
			Message: "Invalid request body",
		})
		return
	}

	err := u.userService.CreateNewUser(registerDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())

		c.JSON(err.Code, responseDTO.ErrorResponse{
			Message: err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "Success",
	})
}

// Profile
// @Summary      get user profile
// @Description  get string by ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /profile [get]
func (u *UserController) Profile(c *gin.Context) {

	//if err != nil {
	//	fmt.Println("error occurred: " + err.Error())
	//	c.JSON(err.Code, err)
	//	return
	//}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})

}

//TODO:Изменение данных профиля
