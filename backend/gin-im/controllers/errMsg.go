package controllers

type ErrMsg struct {
}

var errMsg = make(map[string]string)

// 初始化错误信息映射
func init() {
	errMsg["400001"] = "缺少必要的参数或参数格式错误."
	errMsg["400002"] = "用户名或密码错误,请重新输入."
}

func (m *ErrMsg) String(code string) string {
	return errMsg[code]
}
