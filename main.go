package main

import (
	"go-serve/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	user := controllers.UserController{}
	user.InitUserController(router)

	// router.GET("/users/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	user := findUserById(id, &users)
	// 	if user == nil {
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"message": "No user found",
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, user)
	// })

	// router.POST("/users", func(c *gin.Context) {
	// 	var newUser User
	// 	if err := c.BindJSON(&newUser); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "Could not create user",
	// 		})
	// 		return
	// 	}
	// 	users = append(users, newUser)
	// 	c.JSON(http.StatusCreated, newUser)
	// })

	// router.PUT("/users/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	var newData User
	// 	if err := c.BindJSON(&newData); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "Could not update",
	// 		})
	// 		return
	// 	}
	// 	user := findUserById(id, &users)
	// 	if user == nil {
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"message": "No user found to update",
	// 		})
	// 		return
	// 	}
	// 	user.Name = newData.Name
	// 	user.Username = newData.Username
	// 	user.Email = newData.Email
	// 	user.Password = newData.Password
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Data updated successfully",
	// 	})
	// })

	// router.DELETE("/users", func (c *gin.Context) {
	// 	id := c.Query("id")
	// 	if existingUser := findUserById(id, &users); existingUser == nil {
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"message": "User not found",
	// 		})
	// 		return
	// 	}
	// 	newUsers := []User{}
	// 	for i := range users {
	// 		if users[i].Id != id {
	// 			newUsers = append(newUsers, users[i])
	// 		}
	// 	}
	// 	users = newUsers
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "User deleted successfully",
	// 	})
	// })

	router.Run("localhost:8080")
}
