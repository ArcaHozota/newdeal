package routers

import (
	"fmt"
	"net/http"
	"newdeal/common"
	"newdeal/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("nasbWebToken")

func HomeHandlerInit(r *gin.Engine) {

	count, _ := service.CountHymnsAll()

	homeRouter1 := r.Group("/")
	{
		homeRouter1.GET(common.EmptyString, func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index2.html", gin.H{
				"totalRecords": count,
			})
		})
		homeRouter1.GET("index.action", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index2.html", gin.H{
				"totalRecords": count,
			})
		})
	}

	homeRouter2 := r.Group("/home")
	{
		homeRouter2.GET("index", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index2.html", gin.H{
				"totalRecords": count,
			})
		})
		homeRouter2.GET("page", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index2.html", gin.H{
				"totalRecords": count,
			})
		})
		homeRouter2.GET("to-home-page", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index2.html", gin.H{
				"totalRecords": count,
			})
		})
		homeRouter2.GET("to-inner", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"totalRecords": count,
			})
		})
	}

}

// authMiddleware JWT認証ミドルウェア
func authMiddleware(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/category/not-login")
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
		ctx.Redirect(http.StatusSeeOther, "/category/not-login")
		ctx.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Set("username", claims["username"])
		ctx.Set("loginId", claims["loginId"])
		ctx.Next()
	} else {
		ctx.Redirect(http.StatusSeeOther, "/category/not-login")
		ctx.Abort()
	}
}
