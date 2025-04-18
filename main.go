package main

import (
	"newdeal/common/tools"
	"newdeal/routers"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {

	// Ginを配置する
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"Substr": tools.Substr,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	routers.HomeHandlerInit(r)
	routers.CategoryHandlerInit(r)
	routers.HymnsHandlerInit(r)

	r.Run(":8277")

}
