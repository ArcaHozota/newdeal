package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"newdeal/common"
	"newdeal/common/tools"
	"newdeal/ent"
	"newdeal/routers"
	"newdeal/service"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var tplFS embed.FS

//go:embed static/*
var staticFS embed.FS

func main() {

	// ENTを初期化する
	service.InitEntClient()
	defer func(entClient *ent.Client) {
		err := entClient.Close()
		if err != nil {
			log.Fatalf("failed to close ent core: %v", err)
		}
	}(service.EntClient)

	/* ---------- ①: Ginを配置する ---------- */
	// 本番なら ReleaseMode 推奨
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"Substr": tools.Substr,
	})

	/* ---------- ②: テンプレートを embed から読み込む ---------- */
	tpl := template.Must(template.New(common.EmptyString).ParseFS(tplFS, "templates/*"))
	r.SetHTMLTemplate(tpl)

	/* ---------- ③: 静的ファイルを embed から配信 ---------- */
	// staticFS は "static/css/style.css" という prefix 付きなので、
	// サブ FS に切って URL と一致させる
	subStatic, _ := fs.Sub(staticFS, "static")
	r.StaticFS("/static", http.FS(subStatic))

	// ハンドラを配置する
	routers.BooksHandlerInit(r)
	routers.CategoryHandlerInit(r)
	routers.HomeHandlerInit(r)
	routers.HymnsHandlerInit(r)
	routers.StudentsHandlerInit(r)

	// ポートを定義する
	err := r.Run(":8277")
	if err != nil {
		log.Fatalf("failed to run router: %v", err)
	}

}
