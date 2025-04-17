package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CategoryHandlerInit(r *gin.Engine) {
	categoryRouter := r.Group("/category")
	{
		categoryRouter.GET("/login.action", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "logintoroku.html", nil)
		})
	}
}
