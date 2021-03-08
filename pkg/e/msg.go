package e

//MsgFlags is msg
var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	INVALIDPARAMS:              "请求参数错误",
	ERROREXISTTAG:              "已存在该标签名称",
	ERRORNOTEXISTARTICLE:       "该标签不存在",
	ERRORAUTHCHECKTOKENFAIL:    "Token鉴权失败",
	ERRORAUTHCHECKTOKENTIMEOUT: "Token已超时",
	ERRORAUTHTOKEN:             "Token生成失败",
	ERRORAUTH:                  "Token错误",
}

//GetMsg 根据code获取msg
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
