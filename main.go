package main

import (
	"newdeal/models"
	"newdeal/routers"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"Substr": models.Substr,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	routers.CategoryHandlerInit(r)

	r.Run(":8277")
}
