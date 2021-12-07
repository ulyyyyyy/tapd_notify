package health

import (
	"github.com/gin-gonic/gin"
	"github.com/ulyyyyyy/tapd_notify/internal/helper/ginresp"
)

func IsHealthy(c *gin.Context) {
	ginresp.NewSuccess(c, "OK")
	return
}
