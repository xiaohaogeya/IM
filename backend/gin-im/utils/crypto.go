package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"gin-im/conf"
)

type Crypto struct {
}

// Encode256 sha256加密
func (c *Crypto) Encode256(value string) string {
	secretKey := conf.AppConfig.SecretKey
	m := sha256.New()
	m.Write([]byte(secretKey + value))
	return hex.EncodeToString(m.Sum(nil))
}
