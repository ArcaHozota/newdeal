package routers

import (
	"github.com/gin-gonic/gin"
	"newdeal/common"
)

func StudentsHandlerInit(r *gin.Engine) {

	studentsRouter := r.Group("/students")
	{
		studentsRouter.GET("/to-edition", func(ctx *gin.Context) {
			studentIdStr := ctx.DefaultQuery("editId", common.EmptyString)

		})
	}
}
