package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_EXIST_CANDIDATE      = 30001
	ERROR_EXIST_CANDIDATE_FAIL = 30002
	ERROR_ADD_CANDIDATE_FAIL   = 30003

	ERROR_EXIST_THEME                    = 40001
	ERROR_EXIST_THEME_FAIL               = 40002
	ERROR_ADD_THEME_FAIL                 = 40003
	ERROR_NOT_EXIST_THEME                = 40004
	ERROR_THEME_GET_CANDIDATE_COUNT_FAIL = 40005
	ERROR_THEME_COUNT_NOT_ENOUGH         = 40006

	ERROR_ADD_USER_FAIL         = 50001
	ERROR_CHECK_USER_VALID_FAIL = 50002
	ERROR_INVALID_USER          = 50003

	ERROR_ADD_VOTE_FAIL        = 60001
	ERROR_VOTE_AGAIN_FAILE     = 60002
	ERROR_GET_VOTE_DETAIL_FAIL = 60003
)
