package errorhandler

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler(ginCtx *gin.Context, status int, err error) {
	ginCtx.Error(err)
	ginCtx.AbortWithStatusJSON(status, gin.H{"status": false, "message": err.Error()})
}
