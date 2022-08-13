package biz

import "github.com/cloudwego/hertz/pkg/protocol/consts"

const (
	// 5xx 10000~10100
	CodeServerError = 10000
	// 400 11001~11999
	CodeParamBindError = 11001
	// 401 12001~12999
	// 403 13001~13999
	// 404 14001~14999
	// 400 BizError 20001~99999
)

var (
	codeZhCNText = map[int]string{
		CodeServerError:    "内部服务器错误",
		CodeParamBindError: "参数信息错误",
	}
	codeEnUsText = map[int]string{
		CodeServerError:    "Internal server error",
		CodeParamBindError: "Parameter error",
	}
)

func CodeText(bizCode int) string {
	return codeZhCNText[bizCode]
}

func Code2HttpCode(bizCode int) int {
	if bizCode == 0 {
		return consts.StatusOK
	}
	if bizCode > 20000 {
		return consts.StatusBadRequest
	}
	if 11000 < bizCode && bizCode < 12000 {
		return consts.StatusBadRequest
	}
	if 12000 < bizCode && bizCode < 13000 {
		return consts.StatusUnauthorized
	}
	if 13000 < bizCode && bizCode < 14000 {
		return consts.StatusForbidden
	}
	if 14000 < bizCode && bizCode < 15000 {
		return consts.StatusNotFound
	}
	return consts.StatusInternalServerError
}
