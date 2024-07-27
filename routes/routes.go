package routes

import (
	v1 "task-management/api"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	tm_route := r.Group("/task-management")
	v1.Route(tm_route)
}
