package errno

type Err struct {
	HttpCode int    `json:"http_code"`
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
}

func (e *Err) Error() string {
	return e.Msg
}

func New400(msg string) error {
	return &Err{
		HttpCode: 400,
		Code:     400,
		Msg:      msg,
	}
}

func New500(msg string) error {
	return &Err{
		HttpCode: 500,
		Code:     500,
		Msg:      msg,
	}
}

func NewErr(httpCode int, code int, msg string) error {
	return &Err{
		HttpCode: httpCode,
		Code:     code,
		Msg:      msg,
	}
}
