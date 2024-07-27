package api

import "github.com/gin-gonic/gin"

func RouteV1(r *gin.RouterGroup) {
	r.GET("/ping", handlePing)
}
