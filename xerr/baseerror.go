package xerr

import (
	"github.com/zeromicro/x/errors"
)

// NewEnsumError 在api中使用 错误处理
func NewEnsumError(errCode int) error {
	return errors.New(errCode, MapErrMsg(errCode))
}
func NewSuccessMsg() BaseMessageResponse {
	return BaseMessageResponse{
		Code:    ResponseCodeSuccessLegacy,
		Message: MapErrMsg(ResponseCodeSuccessLegacy),
	}
}
