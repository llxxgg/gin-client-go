package route

import (
	"github.com/gin-gonic/gin"
	"operator/apis"
)

func InitRoute(r *gin.Engine) {
	r.GET("/ns", apis.GetNamespace)
}
