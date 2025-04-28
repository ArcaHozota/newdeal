package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"newdeal/common"
	"newdeal/service"
	"strconv"
)

func StudentsHandlerInit(r *gin.Engine) {

	studentsRouter := r.Group("/students")
	{
		studentsRouter.GET("/to-edition", func(ctx *gin.Context) {
			studentIdStr := ctx.DefaultQuery("editId", common.EmptyString)
			studentId, err := strconv.Atoi(studentIdStr)
			if err != nil || studentId < 1 {
				log.Println(err)
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
			studentById, err := service.GetStudentById(int64(studentId))
			if err != nil {
				log.Println(err)
				ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
					common.AttrNameException: err.Error(),
				})
				return
			}
			ctx.HTML(http.StatusOK, "students-edition.html", gin.H{
				common.AttrNameEntity: studentById,
			})
		})
	}
}
