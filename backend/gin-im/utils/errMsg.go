package utils

/*
	错误信息
*/

type ErrMsg struct {
}

var errMsg = make(map[string]string)

// 初始化错误信息映射
func init() {
	errMsg["400001"] = "缺少必要的参数或参数格式错误."
	errMsg["400002"] = "用户名或密码错误,请重新输入."
	errMsg["400003"] = "权限认证失败,请前去登录."
	errMsg["400004"] = "该用户名已注册,请重新输入."
	errMsg["400005"] = "注册失败,请重新注册."
	errMsg["400006"] = "手机号码格式错误,请重新输入."
	errMsg["400007"] = "邮箱格式错误,请重新输入."
	errMsg["400008"] = "数据库异常."
	errMsg["400009"] = "数据解析失败."
}

func (m *ErrMsg) String(code string) string {
	return errMsg[code]
}
