package controllers

import (
	"gin-im/utils"
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
	errMsg := utils.ErrMsg{}
	msg := errMsg.String(code)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": msg, "code": code, "data": ""})
}

// UserId 获取UserId
func (c *BaseController) UserId(ctx *gin.Context) uint {
	userId, exists := ctx.Get("userId")
	if !exists {
		c.Error(ctx, "400003")
		return 0
	}
	newUserId, ok := userId.(uint)
	if !ok {
		c.Error(ctx, "400003")
		return 0
	}
	return newUserId
}

// AsController 模拟 Django框架的as_view()方法
//func (c *BaseController) AsController(ctx *gin.Context) {
//	c.ctx = ctx
//	method := ctx.Request.Method
//	rVal := reflect.ValueOf(c)
//	m := rVal.MethodByName(method)
//	m.Call(nil)
//}


