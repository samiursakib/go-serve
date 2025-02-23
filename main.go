package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	users := []User{
		{"skdjde", "Alice Johnson", "alicej", "alice@example.com", "password123"},
		{"eirwdc", "Bob Smith", "bobsmith", "bob@example.com", "securePass!"},
		{"znchwr", "Charlie Brown", "charlieb", "charlie@example.com", "charliePass"},
	}

	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := findUserById(id, users)
		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "No user found",
			})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	router.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not create user",
			})
			return
		}
		users = append(users, newUser)
		c.JSON(http.StatusCreated, newUser)
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var newData User
		if err := c.BindJSON(&newData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not update",
			})
			return
		}
		user := findUserById(id, users)
		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "No user found to update",
			})
			return
		}
		user.Name = newData.Name
		user.Username = newData.Username
		user.Email = newData.Email
		user.Password = newData.Password
		c.JSON(http.StatusOK, gin.H{
			"message": "Data updated successfully",
		})
	})

	router.Run("localhost:8080")
}

func findUserById(id string, users []User) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}
