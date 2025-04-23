package routers

import (
	"net/http"
	"newdeal/common"
	"newdeal/service"

	"github.com/gin-gonic/gin"
)

func CategoryHandlerInit(r *gin.Engine) {

	categoryRouter := r.Group("/category")
	{
		categoryRouter.GET("/login", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "logintoroku.html", gin.H{
				"Title": common.EmptyString,
			})
		})
		categoryRouter.GET("/login-with-error", func(ctx *gin.Context) {
			count, err := service.CountHymnsAll()
			if err != nil {
				ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
					"exception": err,
				})
			}
			ctx.HTML(http.StatusOK, "logintoroku.html", gin.H{
				"totalRecords": count,
				"torokuMsg":    common.LoginMsg,
			})
		})
	}

}
