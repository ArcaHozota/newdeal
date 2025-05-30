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

func StudentsHandlerInit(r *gin.Engine) {

	studentsRouter := r.Group("/students")
	{
		studentsRouter.GET("/to-edition", authMiddleware, func(ctx *gin.Context) {
			studentId, exists := ctx.Get("loginId")
			if !exists {
				ctx.Redirect(http.StatusSeeOther, "/category/login-with-error")
				return
			}
			ctx.HTML(http.StatusOK, "students-edition.html", gin.H{
				"studentId": studentId,
			})
		})
		studentsRouter.GET("/initial", authMiddleware, func(ctx *gin.Context) {
			studentIdStr := ctx.DefaultQuery("editId", common.EmptyString)
			studentId, err := strconv.Atoi(studentIdStr)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			studentById, err := service.GetStudentById(int64(studentId))
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, studentById)
		})
		studentsRouter.PUT("info-update", authMiddleware, func(ctx *gin.Context) {
			var req pojos.StudentDTO
			if err := ctx.ShouldBindJSON(&req); err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			updateInfo, err := service.StudentInfoUpdate(req)
			if err != nil {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, updateInfo)
		})
	}

}
