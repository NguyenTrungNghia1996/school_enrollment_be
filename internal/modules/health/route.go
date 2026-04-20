package health

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	healthGroup := router.Group("/health")
	{
		healthGroup.GET("", Check)
	}
}
