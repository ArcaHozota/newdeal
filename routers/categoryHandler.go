package routers

import (
	"log"
	"net/http"
	"newdeal/common"
	"newdeal/service"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("nasbWebToken")

func CategoryHandlerInit(r *gin.Engine) {

	categoryRouter := r.Group("/category")
	{
		categoryRouter.GET("login", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "logintoroku.html", gin.H{
				"Title": common.EmptyString,
			})
		})
		categoryRouter.GET("login-with-error", func(ctx *gin.Context) {
			count, err := service.CountHymnsAll()
			if err != nil {
				ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
					"exception": err,
				})
			}
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"totalRecords": count,
				"torokuMsg":    common.NeedLoginMsg,
			})
		})
		categoryRouter.POST("do-login", func(ctx *gin.Context) {
			var req service.LoginRequest
			// JSONバインディング（リクエストボディから取得）
			if err := ctx.ShouldBindJSON(&req); err != nil {
				ctx.HTML(http.StatusBadRequest, "logintoroku.html", gin.H{
					"errorMsg": common.LoginFormError,
				})
			}
			loginAccount, err := service.ProcessLogin(req)
			if err != nil {
				ctx.HTML(http.StatusBadRequest, "logintoroku.html", gin.H{
					"errorMsg": err,
				})
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": loginAccount,
				"exp":      time.Now().Add(time.Hour * 3).Unix(), // 有効期限：24時間
			})
			tokenString, err := token.SignedString(jwtSecret)
			if err != nil {
				log.Fatalf("トークン作成失敗%v", err)
			}
			ctx.SetCookie(
				"token",                                 // name
				tokenString,                             // value
				3600*3,                                  // maxAge（秒）
				"/",                                     // path
				"shinjukujunfukuinkyokainasb1995.co.uk", // domain（ローカルなら ""）
				true,                                    // secure（HTTPSのみならtrue）
				true,                                    // httpOnly（JavaScriptからアクセス不可）
			)
			ctx.Redirect(http.StatusFound, "/category/to-mainmenu-with-login")
		})
		categoryRouter.GET("to-mainmenu", AuthMiddleware, func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "mainmenu.html", gin.H{
				"loginMsg": common.EmptyString,
			})
		})
		categoryRouter.GET("to-mainmenu-with-login", AuthMiddleware, func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "mainmenu.html", gin.H{
				"loginMsg": common.LoginedMsg,
			})
		})
	}

}

// AuthMiddleware JWT認証ミドルウェア
func AuthMiddleware(ctx *gin.Context) {
	// Cookieから取得（"token" という名前で保存されている想定）
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "トークンが見つかりません"})
		return
	}
	// トークンのパース
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})
	// クレーム確認
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Set("username", claims["username"])
		ctx.Next()
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "無効なトークン"})
	}
}
