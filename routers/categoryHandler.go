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
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"totalRecords": count,
				"torokuMsg":    common.LoginMsg,
			})
		})
		categoryRouter.POST("/do-login", func(ctx *gin.Context) {
			account := ctx.DefaultQuery("loginAcct", common.EmptyString)
			password := ctx.DefaultQuery("userPswd", common.EmptyString)
			loginProcess, err := service.ProcessLogin(account, password)
			if err != nil {
				ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
					"exception": err,
				})
			}
			ctx.HTML(http.StatusOK, "mainmenu.html", gin.H{
				"loginMsg":    common.LoginMsg2,
				"loginStatus": loginProcess,
			})
		})
		categoryRouter.GET("/to-mainmenu", func(ctx *gin.Context) {
			count, err := service.CountHymnsAll()
			if err != nil {
				ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
					"exception": err,
				})
			}
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"totalRecords": count,
				"torokuMsg":    common.LoginMsg,
			})
		})
	}

}
