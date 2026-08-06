package controllers

import (
	"fmt"
	"log"
	"net/http"
	"pm-with-essence/cmd/api/model/userDTO"
	"pm-with-essence/cmd/service/authService"
	"pm-with-essence/cmd/service/userService"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *userService.UserService
	authService *authService.AuthService
}

func NewUserController(userService *userService.UserService, authService *authService.AuthService) *UserController {
	return &UserController{
		userService: userService,
		authService: authService,
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
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := u.userService.CreateNewUser(registerDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// Login
// @Summary login
// @Tags User
// @Description user login
// @Produce json
// @Param loginDTO body userDTO.LoginDTO true "LoginDTO"
// @Success 200 {object} responseDTO.LoginResponse
// @Router       /login [post]
func (u *UserController) Login(c *gin.Context) {
	var loginDTO userDTO.LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := u.authService.Login(loginDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	//TODO:Сделать возвращение токена
	c.JSON(http.StatusOK, token)
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
	c.JSON(http.StatusOK, gin.H{"message": "Success"})

}

//TODO:Обновление токена
//TODO:Логаут
//TODO:Изменение данных профиля
