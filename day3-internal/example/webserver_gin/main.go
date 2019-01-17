package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// START 1 OMIT
func main() {
	router := gin.Default()

	router.GET("/", greet)
	router.Run(":8080")
}

func greet(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

// END 1 OMIT

// START 2 OMIT
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func user(c *gin.Context) {
	c.JSON(http.StatusOK, User{
		ID:   123,
		Name: "Asd Asd",
	})
}

// END 2 OMIT
