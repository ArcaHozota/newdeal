package main

import (
	"newdeal/models"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"Substr": models.Substr,
	})
}
