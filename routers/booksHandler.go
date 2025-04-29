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
		bookRouter.GET("get-chapters", authMiddleware, func(ctx *gin.Context) {
			bookId := ctx.DefaultQuery("bookId", common.EmptyString)
			id, err := strconv.Atoi(bookId)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			chapterDtos, err := service.GetChaptersByBookId(int16(id))
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, chapterDtos)
		})
		bookRouter.POST("info-storage", authMiddleware, func(ctx *gin.Context) {
			var req pojos.PhraseDTO
			if err := ctx.ShouldBindJSON(&req); err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			phInfoMsg, err := service.PhraseInfoStorage(req)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, phInfoMsg)
		})
		bookRouter.GET("to-addition", authMiddleware, func(ctx *gin.Context) {
			bookDtos, err := service.GetBooks()
			if err != nil {
				log.Println(err)
				ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
					common.AttrNameException: err.Error(),
				})
				return
			}
			chapterDtos, err := service.GetChaptersByBookId(1)
			if err != nil {
				log.Println(err)
				ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
					common.AttrNameException: err.Error(),
				})
				return
			}
			ctx.HTML(http.StatusOK, "books-addition.html", gin.H{
				"bookDtos":    bookDtos,
				"chapterDtos": chapterDtos,
			})
		})
	}

}
