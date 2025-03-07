package controllers

import (
	"go-serve/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func (userController *UserController) InitUserController(r *gin.Engine) *gin.RouterGroup {
	userController.userService.InitUserService()
	router := r.Group("/users")
	router.GET("", userController.getUsers())
	return router
}

func (userController *UserController) getUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "okay",
		})
	}
}
