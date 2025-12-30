package xerr

type BizError struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Cause error  `json:"cause,omitempty"`
}

func (biz *BizError) Error() string {
	return biz.Msg
}

var (
	// 通用错误
	ErrInvalidParams = &BizError{Code: 10001, Msg: "Invalid Params"}
	ErrBadRequest    = &BizError{Code: 10002, Msg: "Bad Request"}
	ErrInternal      = &BizError{Code: 10003, Msg: "Internal server error"}
	ErrUnanthorized  = &BizError{Code: 10004, Msg: "Unauthorized"}

	// user module
	ErrUserNotFount  = &BizError{Code: 20001, Msg: "User not exists"}
	ErrUsernameTaken = &BizError{Code: 20002, Msg: "Username has Already been Exists"}
	ErrEmailTaken    = &BizError{Code: 20003, Msg: "Email has Already been Exists"}
)

func NewBizError(code int, msg string, cause error) *BizError {
	return &BizError{Code: code, Msg: msg, Cause: cause}
}

func WrapBiz(code int, msg string, err error) error {
	if err == nil {
		return nil
	}

	return &BizError{Code: code, Msg: msg, Cause: err}
}
