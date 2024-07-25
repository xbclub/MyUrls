package xerr

var message map[int]string

func init() {
	message = make(map[int]string)
	message[ResponseCodeSuccessLegacy] = "SUCCESS"
	message[ResponseCodeParamsCheckError] = "short key already exists"
	message[ResponseCodeServerError] = "系统错误"
}

func MapErrMsg(errcode int) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器繁忙,请稍后再试"
	}
}
