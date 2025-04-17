package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HymnsHandlerInit(r *gin.Engine) {
	hymnsRouter := r.Group("/hymns")
	{
		hymnsRouter.GET("pagination.action", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK)
		})
	}
}
