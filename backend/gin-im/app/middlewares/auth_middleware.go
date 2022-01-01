package middlewares

import (
	"gin-im/auth"
	"gin-im/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
	user, err := auth.ValidateToken(authorization)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": err, "code": code, "data": ""})
		return
	}

	ctx.Set("userId", user.ID)
	ctx.Set("userName", user.UserName)
	ctx.Next()
}
