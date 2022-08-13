package biz

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/errors"
)

type BizError struct {
	HttpCode int    // HTTP 状态码
	BizCode  int    // 业务码
	Msg      string // 错误描述
	Alert    bool   // 是否告警通知
}

func Error(bizCode int) *BizError {
	return &BizError{
		HttpCode: Code2HttpCode(bizCode),
		BizCode:  bizCode,
		Msg:      CodeText(bizCode),
		Alert:    false,
	}
}

func (e *BizError) SetAlert() *BizError {
	e.Alert = true
	return e
}

func (e *BizError) SetHTTPCode(httpCode int) *BizError {
	e.HttpCode = httpCode
	return e
}

func (e *BizError) SetBizCode(bizCode int) *BizError {
	e.BizCode = bizCode
	return e
}

func (e *BizError) SetMsg(msg string) *BizError {
	e.Msg = msg
	return e
}

func (e *BizError) Error() string {
	return fmt.Sprintf(
		"HttpCode=%d\nBizCode=%d\nAlert=%t\nMsg=%s",
		e.HttpCode,
		e.BizCode,
		e.Alert,
		e.Msg,
	)
}

func NewParamBindError(err error) *errors.Error {
	return errors.New(err, errors.ErrorTypePublic, Error(CodeParamBindError))
}
