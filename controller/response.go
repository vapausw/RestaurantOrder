package controller

import (
	co "RestaurantOrder/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

func responseErrorWithMsg(ctx *gin.Context, code co.MyCode, errMsg string) {
	rd := &co.ResponseData{
		Code:    code,
		Message: errMsg,
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func responseSuccess(ctx *gin.Context, data interface{}) {
	rd := &co.ResponseData{
		Code:    co.CodeSuccess,
		Message: "success",
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
