package errorf

const defaultCode = 101

type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Message: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Message
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    e.Code,
		Message: e.Message,
	}
}

type CodeErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCodeErrorResponse(msg string) CodeErrorResponse {

	return CodeErrorResponse{
		Code:    defaultCode,
		Message: msg,
	}
}
