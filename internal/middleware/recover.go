package middleware

import (
	"net/http"

	"go_be_enrollment/internal/common"
	"go_be_enrollment/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			logger.Log.Error("Panic recovered", zap.String("error", err))
		} else {
			logger.Log.Error("Panic recovered", zap.Any("error", recovered))
		}
		common.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
		c.Abort()
	})
}
