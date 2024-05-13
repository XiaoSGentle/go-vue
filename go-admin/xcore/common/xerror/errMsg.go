package xerror

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[NO_RERMIT_ERROR] = "暂无权限！"

	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[TOKEN_FORMAT_ERROR] = "Token格式错误"
	message[TOKEN_ERROR] = "Token错误"

	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"

	message[CAPTCHA_KEY_NOT_FOUND_ERROR] = "请完成验证码"
	message[CAPTCHA_VERIFY_ERROR] = "验证码验证失败"

	message[USER_NOT_EXIST_ERROR] = "用户不存在"
	message[USER_STARTUS_ERROR] = "用户状态异常"

	message[CURD_AFFECT_NONE_ERROR] = "未影响行数"
	message[CURD_UPDATE_AFFECT_NONE_ERROR] = "修改失败"
	message[CURD_DATA_EXIST_ERROR] = "数据已存在"
	message[CURD_DATA_NOT_EXIST_ERROR] = "数据不存在"

	message[DICT_NOT_EXIST_ERROR] = "字典不存在"
}

func MapErrMsg(decode uint32) string {
	if msg, ok := message[decode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(decode uint32) bool {
	if _, ok := message[decode]; ok {
		return true
	} else {
		return false
	}
}
