package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"FileCloud/model"
	"FileCloud/pkg/db"
	"FileCloud/pkg/jwt"
	"FileCloud/response"
	"FileCloud/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Recover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// 打印错误堆栈信息
			msg := fmt.Sprintf("panic: %v", r)
			fmt.Println(msg)
			debug.PrintStack()
			// 封装通用json返回
			// c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			// Result.Fail不是本例的重点，因此用下面代码代替
			response.Error(ctx, 500, errorToString(r))
			// 终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			ctx.Abort()
		}
	}()
	// 加载完 defer recover，继续后续接口调用
	ctx.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

func main() {
	port := viper.GetString("server.port")
	Mysqldb := db.InitDB()
	redis := db.InitRedis()
	model.SetAutoMigrate(Mysqldb)
	defer redis.Close()
	defer Mysqldb.Close()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(Recover)
	r = routers.CollectRoute(r)
	panic(r.Run(":" + port))
}

func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	jwt.JwtKey = []byte(viper.GetString("JwtKey"))
}
