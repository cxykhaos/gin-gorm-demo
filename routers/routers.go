// router
package routers

import (
	"net/http"

	"FileCloud/api"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(Cors())
	r.StaticFS("files", http.Dir("./files")) // 这样可以直接在前端查看到图片不用下载http://localhost:8081/upload_Images/2021/11/09/xEFsXLNOcNltmpzlqhZy.png
	r.GET("/")
	v1 := r.Group("/api/v1")
	{
		v1.GET("health/check", api.TestAPI)

	}
	return r
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
