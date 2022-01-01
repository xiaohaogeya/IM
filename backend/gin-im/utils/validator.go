package utils

import "regexp"

type Validator struct {
}

// ValidateMobile 校验手机号
func (v *Validator) ValidateMobile(mobile string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(mobile)
}

// ValidateEmail 校验邮箱
func (v *Validator) ValidateEmail(email string) bool {
	reg := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(email)
}