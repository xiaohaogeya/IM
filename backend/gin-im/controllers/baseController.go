package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

// Success 成功请求默认状态码为200
func (c *BaseController) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "code": "0", "data": data})
}

// Error 失败请求默认状态码为400；并附带有err信息
func (c *BaseController) Error(ctx *gin.Context, code string) {
	errMsg := ErrMsg{}
	msg := errMsg.String(code)
	ctx.JSON(http.StatusBadRequest, gin.H{"msg": msg, "code": code, "data": ""})
}
