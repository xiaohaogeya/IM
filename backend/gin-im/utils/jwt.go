package utils

import (
	"errors"
	"gin-im/conf"
	"gin-im/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	models.User
	jwt.StandardClaims
}

var (
	ExpireAt  time.Duration // 过期30天
	SecretKey string        // 密钥
)

func init() {
	_ = conf.InitConfig()
	ExpireAt = time.Hour * 24 * time.Duration(conf.AppConfig.JWT.ExpireAt) // 过期30天
	SecretKey = conf.AppConfig.SecretKey                                  // 密钥
}

// GenerateToken 获取jwt token
func GenerateToken(user *models.User) (tokenString string, err error) {
	mySigningKey := []byte(SecretKey)
	expireAt := time.Now().Add(time.Second * time.Duration(ExpireAt)).Unix()
	newUser := *user
	claims := MyCustomClaims{
		newUser,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    string(rune(user.Id)),
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
func ValidateToken(tokenString string) (user models.User, err error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
	if err != nil {
		return user, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		user = claims.User
	} else {
		err = errors.New("validate tokenString failed")
	}
	return
}
