package main

import (
	"newdeal/common/tools"
	"newdeal/routers"
	"newdeal/service"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {

	// ENTを初期化する
	service.InitEntClient()
	defer service.EntCore.Close()

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
