package types

func NewResponse(code int, msg string, data any) *ResultResponse {
	return &ResultResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
func NewDefaultResponse() *ResultResponse {
	return NewResponse(100, "操作成功", nil)
}

func NewDataResponse(data any) *ResultResponse {
	return NewResponse(100, "操作成功", data)
}
