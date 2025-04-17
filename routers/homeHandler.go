package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var count int64 = repository.countHymns()

func HomeHandlerInit(r *gin.Engine) {

	homeRouter1 := r.Group("/")
	{
		homeRouter1.GET("", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"totalRecords": count,
			})
		})
		homeRouter1.GET("index.action", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"totalRecords": count,
			})
		})
	}

	homeRouter2 := r.Group("/home")
	{
		homeRouter2.GET("/login.action", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "logintoroku.html", nil)
		})
	}
}
