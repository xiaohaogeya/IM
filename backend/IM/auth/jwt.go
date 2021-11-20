package auth

import (
	"IM/models"
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	models.User
	jwt.StandardClaims
}

var (
	ExpireAt     = time.Hour * 24 * 30                 // 过期30天
	SecretKey, _ = beego.AppConfig.String("SecretKey") // 密钥
)

// GenerateToken 获取jwt token
func GenerateToken(info *models.User) (tokenString string, err error) {
	mySigningKey := []byte(SecretKey)
	expireAt := time.Now().Add(time.Second * time.Duration(ExpireAt)).Unix()
	user := *info
	claims := MyCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    user.Id,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return
}

// ValidateToken 验证jwt token
func ValidateToken(tokenString string) (info models.User, err error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
	if err != nil {
		return info, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		info = claims.User
	} else {
		err = errors.New("validate tokenString failed")
	}
	return
}
