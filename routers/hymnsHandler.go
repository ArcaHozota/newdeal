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
			pageNumStr := ctx.DefaultQuery(common.AttrNamePageNo, "1")
			pageNum, err := strconv.Atoi(pageNumStr)
			if err != nil || pageNum < 1 {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			cnt, err := service.CountHymnsByKeyword(common.EmptyString)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			hymnDtos, err := service.GetHymnsByKeyword(common.EmptyString, pageNum)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			pageInfos, err := pagination.NewPagination(hymnDtos, cnt, pageNum, int(common.DefaultPageSize), 5)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.JSON(http.StatusOK, pageInfos)
		})
		hymnsRouter.GET("common-retrieve", func(ctx *gin.Context) {
			keyword := ctx.DefaultQuery("keyword", common.EmptyString)
			hymnDtos, err := service.GetHymnsRandomFive(keyword)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.JSON(http.StatusOK, hymnDtos)
		})
		hymnsRouter.GET("kanumi-retrieve", func(ctx *gin.Context) {
			hymnId := ctx.DefaultQuery("hymnId", common.EmptyString)
			id, err := strconv.Atoi(hymnId)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			hymnDtos, err := service.GetHymnsKanumi(int64(id))
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.JSON(http.StatusOK, hymnDtos)
		})
		hymnsRouter.GET("get-info-id", func(ctx *gin.Context) {
			hymnIdStr := ctx.DefaultQuery("hymnId", common.EmptyString)
			hymnId, err := strconv.Atoi(hymnIdStr)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			hymnDto, err := service.GetHymnById(int64(hymnId))
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.JSON(http.StatusOK, hymnDto)
		})
		hymnsRouter.GET("to-pages", authMiddleware, func(ctx *gin.Context) {
			pageNumStr := ctx.DefaultQuery(common.AttrNamePageNo, "1")
			pageNum, err := strconv.Atoi(pageNumStr)
			if err != nil || pageNum < 1 {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.HTML(http.StatusOK, "hymns-pagination.html", gin.H{
				common.AttrNamePageNo: pageNum,
			})
		})
		hymnsRouter.GET("to-addition", authMiddleware, func(ctx *gin.Context) {
			pageNumStr := ctx.DefaultQuery(common.AttrNamePageNo, "1")
			pageNum, err := strconv.Atoi(pageNumStr)
			if err != nil || pageNum < 1 {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.HTML(http.StatusOK, "hymns-addition.html", gin.H{
				common.AttrNamePageNo: pageNum,
			})
		})
		hymnsRouter.GET("to-edition", authMiddleware, func(ctx *gin.Context) {
			pageNumStr := ctx.DefaultQuery(common.AttrNamePageNo, "1")
			pageNum, err := strconv.Atoi(pageNumStr)
			if err != nil || pageNum < 1 {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			editId := ctx.DefaultQuery("editId", common.EmptyString)
			hymnId, err := strconv.Atoi(editId)
			if err != nil {
				log.Println(err)
				ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
					common.AttrNameException: err.Error(),
				})
				return
			}
			hymnDto, err := service.GetHymnById(int64(hymnId))
			if err != nil {
				log.Println(err)
				ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
					common.AttrNameException: err.Error(),
				})
				return
			}
			ctx.HTML(http.StatusOK, "hymns-edition.html", gin.H{
				common.AttrNamePageNo: pageNum,
				common.AttrNameEntity: hymnDto,
			})
		})
		hymnsRouter.GET("to-score-upload", authMiddleware, func(ctx *gin.Context) {
			pageNumStr := ctx.DefaultQuery(common.AttrNamePageNo, "1")
			pageNum, err := strconv.Atoi(pageNumStr)
			if err != nil || pageNum < 1 {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			scoreId := ctx.DefaultQuery("scoreId", common.EmptyString)
			ctx.HTML(http.StatusOK, "hymns-pagination.html", gin.H{
				common.AttrNamePageNo: pageNum,
				"scoreId":             scoreId,
			})
		})
		hymnsRouter.POST("score-upload", authMiddleware, func(ctx *gin.Context) {
			var req pojos.HymnDTO
			if err := ctx.ShouldBindJSON(&req); err != nil {
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			storageInfo, err := service.HymnScoreStorage(req)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.JSON(http.StatusOK, storageInfo)
		})
		hymnsRouter.PUT("info-update", authMiddleware, func(ctx *gin.Context) {
			var req pojos.HymnDTO
			if err := ctx.ShouldBindJSON(&req); err != nil {
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			updateInfo, err := service.HymnInfoUpdate(req)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.JSON(http.StatusOK, updateInfo)
		})
		hymnsRouter.GET("to-random-five", authMiddleware, func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "hymns-random-five.html", gin.H{})
		})
	}

}
