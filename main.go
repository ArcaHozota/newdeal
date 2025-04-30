package main

import (
	"log"
	"newdeal/common/tools"
	"newdeal/ent"
	"newdeal/routers"
	"newdeal/service"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {

	// ENTを初期化する
	service.InitEntClient()
	defer func(entCore *ent.Client) {
		err := entCore.Close()
		if err != nil {
			log.Fatalf("failed to close ent core: %v", err)
		}
	}(service.EntCore)

	// Ginを配置する
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"Substr": tools.Substr,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// ハンドラを配置する
	routers.BookHandlerInit(r)
	routers.CategoryHandlerInit(r)
	routers.HymnsHandlerInit(r)
	routers.HomeHandlerInit(r)
	routers.StudentsHandlerInit(r)

	// ポートを定義する
	err := r.Run(":8277")
	if err != nil {
		log.Fatalf("failed to run router: %v", err)
	}

}
