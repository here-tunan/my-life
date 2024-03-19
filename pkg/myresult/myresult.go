package myresult

type MyResult[T any] struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Data    T      `json:"data"`
}

func (result *MyResult[T]) OK(data T) *MyResult[T] {
	result.Success = true
	result.Code = 200
	result.Data = data
	return result
}

func (result *MyResult[T]) Err(code int, msg string) *MyResult[T] {
	result.Success = false
	result.Code = code
	result.Error = msg
	return result
}
