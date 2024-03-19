package myerror

type MyError struct {
	msg string
}

// 实现  Error() 接口
func (err MyError) Error() string {
	return err.msg
}

func NewError(code Code) error {
	return MyError{msg: code.String()}
}

func NewErrorMsg(msg string) error {
	return MyError{msg: msg}
}
