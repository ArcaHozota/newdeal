package routers

import (
	"log"
	"net/http"
	"newdeal/common"
	"newdeal/common/pagination"
	"newdeal/pojos"
	"newdeal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HymnsHandlerInit(r *gin.Engine) {

	hymnsRouter := r.Group("/hymns")
	{
		hymnsRouter.GET("pagination", func(ctx *gin.Context) {
			// 未指定時は 1 にしたいので DefaultQuery
			pageNumStr := ctx.DefaultQuery("pageNum", "1")
			pageNum, err := strconv.Atoi(pageNumStr)
			if err != nil || pageNum < 1 {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			cnt, err := service.CountHymnsByKeyword(common.EmptyString)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			hymnDtos, err := service.GetHymnsByKeyword(common.EmptyString, pageNum)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			pageInfos, err := pagination.NewPagination(hymnDtos, cnt, pageNum, int(common.DefaultPageSize), 5)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			ctx.JSON(http.StatusOK, pageInfos)
		})
		hymnsRouter.GET("common-retrieve", func(ctx *gin.Context) {
			keyword := ctx.DefaultQuery("keyword", common.EmptyString)
			hymnDtos, err := service.GetHymnsRandomFive(keyword)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			ctx.JSON(http.StatusOK, hymnDtos)
		})
		hymnsRouter.GET("kanumi-retrieve", func(ctx *gin.Context) {
			hymnId := ctx.DefaultQuery("hymnId", common.EmptyString)
			id, err := strconv.Atoi(hymnId)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			hymnDtos, err := service.GetHymnsKanumi(int64(id))
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			ctx.JSON(http.StatusOK, hymnDtos)
		})
		hymnsRouter.GET("get-info-id", func(ctx *gin.Context) {
			keyword := ctx.DefaultQuery("hymnId", common.EmptyString)
			hymnId, err := strconv.ParseInt(keyword, 10, 64)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			hymn, err := service.GetHymnById(hymnId)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
			}
			hymnDto := pojos.HymnDTO{
				ID:          strconv.FormatInt(hymn.ID, 10),
				NameJP:      hymn.NameJp,
				NameKR:      hymn.NameKr,
				Serif:       hymn.Serif,
				Link:        hymn.Link,
				Score:       nil,
				Biko:        common.EmptyString,
				UpdatedUser: strconv.FormatInt(hymn.UpdatedUser, 10),
				UpdatedTime: hymn.UpdatedTime,
				LineNumber:  pojos.LineNumber(5),
			}
			ctx.JSON(http.StatusOK, hymnDto)
		})
	}

}
