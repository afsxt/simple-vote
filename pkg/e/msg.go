package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	ERROR_EXIST_CANDIDATE:      "该候选人已经存在",
	ERROR_EXIST_CANDIDATE_FAIL: "检查候选人是否存在失败",
	ERROR_ADD_CANDIDATE_FAIL:   "新增候选人失败",

	ERROR_EXIST_THEME:                    "该主题已经存在",
	ERROR_EXIST_THEME_FAIL:               "检查主题是否存在失败",
	ERROR_NOT_EXIST_THEME:                "该主题不存在",
	ERROR_ADD_THEME_FAIL:                 "新增主题失败",
	ERROR_THEME_GET_CANDIDATE_COUNT_FAIL: "获取选举主题候选人总数失败",
	ERROR_THEME_COUNT_NOT_ENOUGH:         "该主题候选人不够",

	ERROR_ADD_USER_FAIL:         "新增用户失败",
	ERROR_CHECK_USER_VALID_FAIL: "较验用户是否合法失败",
	ERROR_INVALID_USER:          "非法用户",

	ERROR_ADD_VOTE_FAIL:        "新增投票失败",
	ERROR_VOTE_AGAIN_FAILE:     "该用户已经对该主题投过票",
	ERROR_GET_VOTE_DETAIL_FAIL: "用户获取投票详情失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
