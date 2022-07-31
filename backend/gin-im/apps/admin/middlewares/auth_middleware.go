package middlewares

import (
	"gin-im/apps/admin/auth"
	"gin-im/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
}

// CheckAuth 校验jwt认证
func (m *AuthMiddleware) CheckAuth(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")
	code := "400003"
	errMsg := utils.ErrMsg{}
	msg := errMsg.String(code)

	if authorization == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": msg, "code": code, "data": ""})
		return
	}
	userId, username, err := auth.ValidateToken(authorization)
	if err != nil || userId == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": msg, "code": code, "data": ""})
		return
	}
	ctx.Set("userId", userId)
	ctx.Set("userName", username)
	ctx.Next()
}
