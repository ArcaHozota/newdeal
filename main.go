package main

import (
	"log"
	"newdeal/common"
	"newdeal/routers"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"Substr": common.Substr,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	routers.HomeHandlerInit(r)
	routers.CategoryHandlerInit(r)
	routers.HymnsHandlerInit(r)

	err := r.Run(":8277")
	if err != nil {
		log.Println(err.Error())
		return
	}

}
