package routers

import (
	"net/http"
	"newdeal/common"
	"newdeal/repository"

	"github.com/gin-gonic/gin"
)

var count uint32 = repository.CountHymns()

func HomeHandlerInit(r *gin.Engine) {

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
		homeRouter2.GET("index.action", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"totalRecords": count,
			})
		})
		homeRouter2.GET("page.action", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"totalRecords": count,
			})
		})
		homeRouter2.GET("toHomePage.action", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"totalRecords": count,
			})
		})
		homeRouter2.GET("toIchiranhyo.action", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index2.html", gin.H{
				"totalRecords": count,
			})
		})
	}
}
