package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorlazzaroli/microservicesTest/auth/config"
	"github.com/victorlazzaroli/microservicesTest/auth/services"
	"gorm.io/gorm"
)

type usersController struct {
	service services.UsersServicesI
}

type UsersControllerI interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)
	Signout(c *gin.Context)
}

func (u *usersController) Signup(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})

		return
	}

	err := u.service.Signup(body.Name, body.Email, body.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{})
}

// Signin implements UsersControllerI.
func (u *usersController) Signin(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})

		return
	}

	token, err := u.service.Signin(body.Email, body.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	log.Println("token: ", token)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

// Signout implements UsersControllerI.
func (u *usersController) Signout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
}

func NewUsersController(engine *gin.Engine, db *gorm.DB, env *config.Env) UsersControllerI {
	service := services.NewUsersService(db, env)
	controller := &usersController{
		service: service,
	}
	api := engine.Group("auth")
	{
		api.POST("signup", controller.Signup)
		api.POST("signin", controller.Signin)
		api.POST("signout", controller.Signout)
	}

	return controller
}
