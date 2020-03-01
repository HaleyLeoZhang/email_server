package e

var MsgFlags = map[int]string{
	SUCCESS: "success",
	ERROR:   "服务器繁忙",

	INVALID_PARAMS:   "请求参数错误",
	SOURCE_NOT_FOUND: "找不到指定资源",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
