package routers

import (
	"log"
	"net/http"
	"newdeal/common"
	"newdeal/common/pagination"
	"newdeal/service"

	"github.com/gin-gonic/gin"
)

func HymnsHandlerInit(r *gin.Engine) {

	hymnsRouter := r.Group("/hymns")
	{
		hymnsRouter.GET("pagination.action", func(ctx *gin.Context) {
			pageNum := ctx.GetInt("pageNum")
			cnt, err := service.CountHymnsByKeyword(common.EmptyString)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			dtos, err := service.GetHymnsByKeyword(common.EmptyString, pageNum)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			pageInfos, err := pagination.NewPagination(dtos, cnt, pageNum, int(common.DefaultPageSize), 5)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			ctx.JSON(http.StatusOK, pageInfos)
		})
	}

}
