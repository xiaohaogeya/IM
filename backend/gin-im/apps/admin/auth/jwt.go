package auth

import (
	"errors"
	"gin-im/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	UserID   uint
	UserName string
	jwt.StandardClaims
}

var (
	ExpireAt  time.Duration // 过期30天
	SecretKey string        // 密钥
	Issuer    string        // 签发者
)

func init() {
	ExpireAt = time.Hour * 24 * time.Duration(conf.AppConfig.JWT.ExpireAt) // 默认过期30天
	Issuer = conf.AppConfig.JWT.Issuer
	SecretKey = conf.AppConfig.SecretKey
}

// GenerateToken 获取jwt token
func GenerateToken(userId uint, username string) (tokenString string, err error) {
	mySigningKey := []byte(SecretKey)
	expireAt := time.Now().Add(ExpireAt).Unix()
	claims := MyCustomClaims{
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    Issuer,
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
func ValidateToken(tokenString string) (userId uint, username string, err error) {
	token, _ := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		userId = claims.UserID
		username = claims.UserName
	} else {
		err = errors.New("validate tokenString failed")
	}
	return
}

//func ValidateToken(tokenString string) (user models.User, err error) {
//	token, err := jwt.ParseWithClaims(
//		tokenString,
//		&MyCustomClaims{},
//		func(token *jwt.Token) (interface{}, error) {
//			return []byte(SecretKey), nil
//		})
//	if err != nil {
//		return user, err
//	}
//	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
//		user = claims.User
//	} else {
//		err = errors.New("validate tokenString failed")
//	}
//	return
//}
