package api

import (
	"FileCloud/response"

	"github.com/gin-gonic/gin"
)

func TestAPI(ctx *gin.Context) {
	response.MsgSuccess(ctx, "success")
}
