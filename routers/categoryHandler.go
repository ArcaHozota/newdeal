package routers

import (
	"fmt"
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
			ctx.HTML(http.StatusOK, "login-toroku.html", gin.H{
				common.AttrNameTorokuMsg: common.EmptyString,
			})
		})
		categoryRouter.GET("login-with-out", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "login-toroku.html", gin.H{
				common.AttrNameTorokuMsg: common.LogoutMsg,
			})
		})
		categoryRouter.GET("login-with-error", func(ctx *gin.Context) {
			count, err := service.CountHymnsAll()
			if err != nil {
				log.Println(err)
				ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
					common.AttrNameException: err.Error(),
				})
				return
			}
			ctx.HTML(http.StatusUnauthorized, "index.html", gin.H{
				"totalRecords":           count,
				common.AttrNameTorokuMsg: common.NeedLoginMsg,
			})
		})
		categoryRouter.POST("do-login", func(ctx *gin.Context) {
			var req service.LoginRequest
			// JSONバインディング（リクエストボディから取得）
			if err := ctx.ShouldBind(&req); err != nil {
				log.Println(err)
				ctx.HTML(http.StatusBadRequest, "login-toroku.html", gin.H{
					common.AttrNameTorokuMsg: common.LoginFormError,
				})
				return
			}
			loginUser, err := service.ProcessLogin(req)
			if err != nil {
				log.Println(err)
				ctx.HTML(http.StatusBadRequest, "login-toroku.html", gin.H{
					common.AttrNameTorokuMsg: err.Error(),
				})
				return
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": loginUser.Username,
				"loginId":  loginUser.ID,
				"exp":      time.Now().Add(time.Hour * 3).Unix(), // 有効期限：24時間
			})
			tokenString, err := token.SignedString(jwtSecret)
			if err != nil {
				log.Fatalf("トークン作成失敗%v", err)
			}
			ctx.SetCookie(
				"token",     // name
				tokenString, // value
				3600*3,      // maxAge（秒）
				"/",         // path
				"",          // domain（ローカルなら ""）shinjukujunfukuinkyokainasb1995.co.uk
				true,        // secure（HTTPSのみならtrue）
				true,        // httpOnly（JavaScriptからアクセス不可）
			)
			ctx.Redirect(http.StatusSeeOther, "/category/to-mainmenu-with-login")
		})
		categoryRouter.POST("do-logout", authMiddleware, func(ctx *gin.Context) {
			// Cookieを削除（有効期限を過去に設定）
			ctx.SetCookie(
				"token", // Cookie名
				"",      // 空にする
				-1,      // 負数の maxAge で削除
				"/",     // パス
				"",      // ドメイン（ローカルなら空でOK）shinjukujunfukuinkyokainasb1995.co.uk
				true,    // Secure
				true,    // HttpOnly
			)
			ctx.Redirect(http.StatusSeeOther, "/category/login-with-out") // ログアウト後ログインページへ
		})
		categoryRouter.GET("to-mainmenu", authMiddleware, func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "mainmenu.html", gin.H{
				common.AttrNameTorokuMsg: common.EmptyString,
			})
		})
		categoryRouter.GET("to-mainmenu-with-login", authMiddleware, func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "mainmenu.html", gin.H{
				common.AttrNameTorokuMsg: common.LoginedMsg,
			})
		})
		categoryRouter.GET("get-username", authMiddleware, func(ctx *gin.Context) {
			username, exists := ctx.Get("username")
			if !exists {
				ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
					common.AttrNameException: common.StudentError,
				})
			}
			ctx.JSON(http.StatusOK, username)
		})
	}

}

// authMiddleware JWT認証ミドルウェア
func authMiddleware(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/category/login-with-error")
		ctx.Abort()
		return
	}
	// トークンのパースと署名方式チェック
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// algがHS256かチェック
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		ctx.Redirect(http.StatusSeeOther, "/category/login-with-error")
		ctx.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Set("username", claims["username"])
		ctx.Set("loginId", claims["loginId"])
		ctx.Next()
	} else {
		ctx.Redirect(http.StatusSeeOther, "/category/login-with-error")
		ctx.Abort()
	}
}
