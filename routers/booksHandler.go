package routers

import (
	"log"
	"net/http"
	"newdeal/common"
	"newdeal/pojos"
	"newdeal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookHandlerInit(r *gin.Engine) {

	bookRouter := r.Group("/books")
	{
		bookRouter.GET("get-chapters", func(ctx *gin.Context) {
			bookId := ctx.DefaultQuery("bookId", common.EmptyString)
			id, err := strconv.Atoi(bookId)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			chapterDtos, err := service.GetChaptersByBookId(int16(id))
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.JSON(http.StatusOK, chapterDtos)
		})
		bookRouter.POST("info-storage", func(ctx *gin.Context) {
			var req pojos.PhraseDTO
			// JSONバインディング（リクエストボディから取得）
			if err := ctx.ShouldBindJSON(&req); err != nil {
				ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
					common.AttrNameException: err,
				})
				return
			}
			phInfoMsg, err := service.PhraseInfoStorage(req)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			ctx.JSON(http.StatusOK, phInfoMsg)
		})
		bookRouter.GET("to-addition", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "books-addition.html", gin.H{})
		})
	}

}
