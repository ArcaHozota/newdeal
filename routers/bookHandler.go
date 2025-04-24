package routers

import (
	"log"
	"net/http"
	"newdeal/common"
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
	}
}
