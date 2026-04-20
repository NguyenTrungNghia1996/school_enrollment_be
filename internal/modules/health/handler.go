package health

import (
	"net/http"

	"go_be_enrollment/internal/common"

	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	common.SuccessResponse(c, http.StatusOK, "Service is up and running", gin.H{
		"status": "OK",
	})
}
