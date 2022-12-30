package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func response(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	res := gin.H{
		"code": code,
		"msg":  msg,
	}
	if data != nil {
		res["data"] = data
	}
	ctx.JSON(httpStatus, res)
}

func success(ctx *gin.Context, data interface{}, msg string) {
	response(ctx, http.StatusOK, 200, data, msg)
}

func MsgSuccess(ctx *gin.Context, msg string) {
	success(ctx, nil, msg)
}

func DataSuccess(ctx *gin.Context, data interface{}) {
	success(ctx, data, "success")
}

func Error(ctx *gin.Context, code int, msg string) {
	response(ctx, http.StatusInternalServerError, code, nil, msg)
}

func fail(ctx *gin.Context, code int, data interface{}, msg string) {
	response(ctx, code, code, data, msg)
}

func MsgFail(ctx *gin.Context, code int, msg string) {
	fail(ctx, code, nil, msg)
}

func DataFail(ctx *gin.Context, code int, data interface{}) {
	fail(ctx, code, data, "fail")
}

func Redirect(ctx *gin.Context, page string) {
	ctx.Redirect(http.StatusFound, page)
}
