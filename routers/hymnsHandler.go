package routers

import (
	"net/http"
	"newdeal/common"
	"newdeal/repository"

	"github.com/gin-gonic/gin"
)

func HymnsHandlerInit(r *gin.Engine) {
	hymnsRouter := r.Group("/hymns")
	{
		hymnsRouter.GET("pagination.action", func(ctx *gin.Context) {
			pageNum := ctx.GetUint("pageNum")
			offset := uint32(common.DefaultPageSize) * (uint32(pageNum) - 1)
			hymns := repository.PaginationHymns(common.EmptyString, common.DefaultPageSize, offset)
			ctx.JSON(http.StatusOK, hymns)
		})
	}
}
