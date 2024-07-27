package api

import (
	api_v1 "task-management/api/v1"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	api_v1.RouteV1(v1)
}
