package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luishsr/go-p2p-lending/internal/routes"
)

func main() {
	router := gin.Default()

	router.POST("/login", routes.LoginHandler) // Setup the login route

	router.Run(":8080")
}
