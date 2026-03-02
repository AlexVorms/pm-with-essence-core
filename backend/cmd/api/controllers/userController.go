package controllers

import (
	"log"
	"net/http"
	"pm-with-essence/cmd/api/model/userDTO"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
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
	//if err != nil {
	//	fmt.Println("error occurred: " + err.Error())
	//	c.JSON(err.Code, err)
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// Login
// @Summary login
// @Tags User
// @Description user login
// @Produce json
// @Param loginDTO body userDTO.LoginDTO true "LoginDTO"
// @Success 200
// @Router       /login [post]
func (u *UserController) Login(c *gin.Context) {
	var loginDTO userDTO.LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//if err != nil {
	//	fmt.Println("error occurred: " + err.Error())
	//	c.JSON(err.Code, err)
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
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
