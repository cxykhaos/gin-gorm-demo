package middle

import (
	"net/http"
	"time"

	"FileCloud/pkg/jwt"
	"FileCloud/response"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(ctx *gin.Context) {
	tokenString, _ := ctx.Cookie("token")
	if tokenString == "" {
		response.MsgFail(ctx, http.StatusUnauthorized, "权限不够")
		ctx.Abort()
		return
	}
	token, claims, err := jwt.ParseToken(tokenString)
	if time.Now().Unix() > claims.ExpiresAt {
		response.MsgFail(ctx, http.StatusUnauthorized, "token过期")
		ctx.Abort()
		return
	}
	if err != nil || !token.Valid {
		response.MsgFail(ctx, http.StatusUnauthorized, "权限不够1")
		ctx.Abort()
		return
	}
	// userId := claims.User_Id
	// user := model.User{User_Id: userId}
	// user.User_Info()
	// if user.User_Id == "" {
	// 	response.MsgFail(ctx, http.StatusUnauthorized, "权限不够2")
	// 	ctx.Abort()
	// 	return
	// }
	// Set global key
	// ctx.Set("user", user)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		GetUserProfile(ctx)
		ctx.Next()
	}
}
