package routers

import (
	"net/http"
	"newdeal/common"
	"newdeal/service"

	"github.com/gin-gonic/gin"
)

func HymnsHandlerInit(r *gin.Engine) {
	hymnsRouter := r.Group("/hymns")
	{
		hymnsRouter.GET("pagination.action", func(ctx *gin.Context) {
			pageNum := ctx.GetUint("pageNum")
			dtos := service.GetHymnsByKeyword(common.EmptyString, uint32(pageNum))
			ctx.JSON(http.StatusOK, dtos)
		})
	}
}
