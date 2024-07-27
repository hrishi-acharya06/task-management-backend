package main

import (
	"os"
	"task-management/constants"
	"task-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	routes.Route(r)
	r.Run(":" + os.Getenv(constants.ServerPort))
}
