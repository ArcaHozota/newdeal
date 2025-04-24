package routers

import (
	"net/http"
	"newdeal/common"
	"newdeal/service"

	"github.com/gin-gonic/gin"
)

func HomeHandlerInit(r *gin.Engine) {

	count, _ := service.CountHymnsAll()

	homeRouter1 := r.Group("/")
	{
		homeRouter1.GET(common.EmptyString, func(ctx *gin.Context) {
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
		homeRouter2.GET("index", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"totalRecords": count,
			})
		})
		homeRouter2.GET("page", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"totalRecords": count,
			})
		})
		homeRouter2.GET("to-home-page", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"totalRecords": count,
			})
		})
		homeRouter2.GET("to-list", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index2.html", gin.H{
				"totalRecords": count,
			})
		})
	}

}
