package wgin

type ReturnCode struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

var SuccessCode = 0

func SetSuccessCode(code int) {
	SuccessCode = code
}
