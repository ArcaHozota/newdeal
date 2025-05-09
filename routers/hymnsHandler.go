package routers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"newdeal/common"
	"newdeal/common/pagination"
	"newdeal/pojos"
	"newdeal/service"
	"strconv"
	"strings"
	"time"

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
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			cnt, err := service.CountHymnsByKeyword(common.EmptyString)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			hymnDtos, err := service.GetHymnsByKeyword(common.EmptyString, pageNum)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			pageInfos, err := pagination.NewPagination(hymnDtos, cnt, pageNum, int(common.DefaultPageSize), 5)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, pageInfos)
		})
		hymnsRouter.GET("common-retrieve", func(ctx *gin.Context) {
			keyword := ctx.DefaultQuery("keyword", common.EmptyString)
			hymnDtos, err := service.GetHymnsRandomFive(keyword)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, hymnDtos)
		})
		hymnsRouter.GET("kanumi-retrieve", func(ctx *gin.Context) {
			log.Println("kanumi-retrieve開始")
			time01 := time.Now()
			hymnId := ctx.DefaultQuery("hymnId", common.EmptyString)
			id, err := strconv.Atoi(hymnId)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.Header("X-Accel-Buffering", "no")
			ctx.Header("Content-Type", "text/event-stream")
			ctx.Header("Cache-Control", "no-cache")
			ctx.Header("Connection", "keep-alive")
			flusher, ok := ctx.Writer.(http.Flusher)
			if !ok {
				ctx.JSON(http.StatusInternalServerError, "Streaming not supported")
				return
			}
			// 每30秒发送一次空数据
			ticker := time.NewTicker(30 * time.Second)
			defer ticker.Stop()
			done := make(chan struct{})
			var hymnDtos []pojos.HymnDTO
			go func() {
				// 長時間サービス
				hymnDtos, err = service.GetHymnsKanumi(int64(id))
				if err != nil {
					log.Println(err)
					ctx.JSON(http.StatusBadRequest, err.Error())
					return
				}
				close(done)
			}()
			for {
				select {
				case <-done:
					// 写一条真正的事件，把 JSON 打包在 data: 行里：
					b, _ := json.Marshal(hymnDtos)
					_, _ = fmt.Fprintf(ctx.Writer, "data: %s\n\n", b)
					flusher.Flush()
					duration := time.Now().Sub(time01).Seconds()
					log.Printf("kanumi-retrieve終了、かかる時間：%v\n", duration)
					return
				case <-ticker.C:
					_, _ = fmt.Fprintf(ctx.Writer, ": keep-alive\n\n")
					// 注释，不影响输出
					flusher.Flush()
				}
			}
		})
		hymnsRouter.GET("get-info-id", func(ctx *gin.Context) {
			hymnIdStr := ctx.DefaultQuery("hymnId", common.EmptyString)
			hymnId, err := strconv.Atoi(hymnIdStr)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			hymnDto, err := service.GetHymnById(int64(hymnId))
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, hymnDto)
		})
		hymnsRouter.GET("get-records", func(ctx *gin.Context) {
			count, err := service.CountHymnsAll()
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, count)
		})
		hymnsRouter.GET("score-download", func(ctx *gin.Context) {
			scoreIdStr := ctx.DefaultQuery("scoreId", common.EmptyString)
			scoreId, err := strconv.Atoi(scoreIdStr)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			download, err := service.HymnScoreDownload(int64(scoreId))
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			biko := download.Biko
			index := strings.Index(biko, "/") + 1
			mimeType := biko[index:]
			nameStr := strconv.Itoa(int(download.WorkID))
			ctx.Header("Content-Type", mimeType)
			ctx.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, nameStr+"."+mimeType))
			ctx.Header("Content-Transfer-Encoding", "binary")
			ctx.Header("Cache-Control", "no-cache")
			ctx.Writer.WriteHeader(http.StatusOK)
			_, _ = ctx.Writer.Write(download.Score)
		})
		hymnsRouter.GET("to-pages", authMiddleware, func(ctx *gin.Context) {
			pageNumStr := ctx.DefaultQuery(common.AttrNamePageNo, "1")
			pageNum, err := strconv.Atoi(pageNumStr)
			if err != nil || pageNum < 1 {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
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
				ctx.JSON(http.StatusBadRequest, err.Error())
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
				ctx.JSON(http.StatusBadRequest, err.Error())
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
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			scoreId := ctx.DefaultQuery("scoreId", common.EmptyString)
			ctx.HTML(http.StatusOK, "hymns-score-upload.html", gin.H{
				common.AttrNamePageNo: pageNum,
				"scoreId":             scoreId,
			})
		})
		hymnsRouter.POST("score-upload", authMiddleware, func(ctx *gin.Context) {
			var req pojos.HymnDTO
			if err := ctx.ShouldBindJSON(&req); err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			storageInfo, err := service.HymnScoreStorage(req)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, storageInfo)
		})
		hymnsRouter.POST("info-storage", authMiddleware, func(ctx *gin.Context) {
			var req pojos.HymnDTO
			if err := ctx.ShouldBindJSON(&req); err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			studentId, exists := ctx.Get("loginId")
			if !exists {
				ctx.Redirect(http.StatusSeeOther, "/category/login-with-error")
				return
			}
			loginIdStr, ok := studentId.(string)
			if !ok {
				ctx.Redirect(http.StatusSeeOther, "/category/login-with-error")
				return
			}
			loginId, err := strconv.Atoi(loginIdStr)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			updateInfo, err := service.HymnInfoStorage(req, int64(loginId))
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, updateInfo)
		})
		hymnsRouter.PUT("info-update", authMiddleware, func(ctx *gin.Context) {
			var req pojos.HymnDTO
			if err := ctx.ShouldBindJSON(&req); err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			studentId, exists := ctx.Get("loginId")
			if !exists {
				ctx.Redirect(http.StatusSeeOther, "/category/login-with-error")
				return
			}
			loginIdStr, ok := studentId.(string)
			if !ok {
				ctx.Redirect(http.StatusSeeOther, "/category/login-with-error")
				return
			}
			loginId, err := strconv.Atoi(loginIdStr)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			updateInfo, err := service.HymnInfoUpdate(req, int64(loginId))
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, updateInfo)
		})
		hymnsRouter.GET("to-random-five", authMiddleware, func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "hymns-random-five.html", gin.H{})
		})
		hymnsRouter.GET("random-five-retrieve", authMiddleware, func(ctx *gin.Context) {
			keyword := ctx.DefaultQuery("keyword", common.EmptyString)
			hymnDtos, err := service.GetHymnsRandomFive(keyword)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, hymnDtos)
		})
		hymnsRouter.GET("deletion-check", authMiddleware, func(ctx *gin.Context) {
			studentId, exists := ctx.Get("loginId")
			if !exists {
				ctx.Redirect(http.StatusSeeOther, "/category/login-with-error")
				return
			}
			loginIdStr, ok := studentId.(string)
			if !ok {
				ctx.Redirect(http.StatusSeeOther, "/category/login-with-error")
				return
			}
			loginId, err := strconv.Atoi(loginIdStr)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			auth, err := service.CheckDeleteAuth(int64(loginId))
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusForbidden, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, auth)
		})
	}

}
