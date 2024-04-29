package apis

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/util/json"
	"operator/service"
)

func GetNamespace(c *gin.Context) {
	ns := service.GetNamespace(c)
	result := map[string]interface{}{
		"ns": ns,
	}
	str, err := json.Marshal(result)
	if err != nil {
		c.String(500, "error")
	}
	c.String(200, string(str))
}