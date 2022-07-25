package projectTest

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	user := struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}{
		UserName: "xiaohao",
		Password: "123456",
	}
	str, err := json.Marshal(&user)
	if err != nil {
		return 
	}
	httptest.NewRequest("post", "/api/user/login", strings.NewReader(string(str)))
}
