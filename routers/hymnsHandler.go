package routers

import (
	"log"
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
			dtos, err := service.GetHymnsByKeyword(common.EmptyString, uint32(pageNum))
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			ctx.JSON(http.StatusOK, dtos)
		})
	}
}
