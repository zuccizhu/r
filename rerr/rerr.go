package rerr

import "net/http"

var (
	BizErr        = newBizError(http.StatusInternalServerError, "BIZ_ERROR", "业务错误")
	UnknownErr    = newBizError(http.StatusInternalServerError, "UNKNOWN_ERROR", "未知错误")
	ParamErr      = newBizError(http.StatusBadRequest, "PARAM_ERROR", "参数错误")
	NotSupportErr = newBizError(http.StatusNotImplemented, "NOT_SUPPORT_ERROR", "不支持的操作")
	DbErr         = newBizError(http.StatusInternalServerError, "DB_ERROR", "数据库错误")
)

type BizError struct {
	Status int32  `json:"status"`
	Code   string `json:"code"`
	Msg    string `json:"msg"`
}

func (e *BizError) Error() string {
	return e.Msg
}

func NewBizErr(code string, msg string) *BizError {
	return newBizError(http.StatusInternalServerError, code, msg)
}

func newBizError(status int32, code string, msg string) *BizError {
	return &BizError{
		Status: status,
		Code:   code,
		Msg:    msg,
	}
}
