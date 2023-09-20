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

func (r ReturnCode) GetCode() int {
	return r.Code
}

func (r ReturnCode) GetMsg() string {
	return r.Msg
}

func (r ReturnCode) GetData() interface{} {
	return r.Data
}

func (r *ReturnCode) SetMsg(msg string) ReturnCode{
	r.Msg = msg
	return *r
}

type IReturnCode interface {
	GetCode() int
	GetMsg() string
	GetData() interface{}
}
