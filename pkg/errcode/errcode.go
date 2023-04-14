package errcode

var (
	Success         = 0
	Failed          = 10000
	ServerError     = 10001
	NotFound        = 10002
	RequestOverflow = 10005

	Unauthorized    = 11003
	InvalidToken    = 11004
	InvalidPassword = 11005
	InvalidUsername = 11006
	HashError       = 11007

	DbQueryFailed  = 20001
	DbExists       = 20002
	DbNotExists    = 20003
	DbCreatFailed  = 20004
	DbUpdateFailed = 20005

	InvalidParam = 21001
)

var errsMap = make(map[int]string)

func init() {
	errsMap[Success] = "成功"
	errsMap[Failed] = "失败"
}

func GetDesc(code int) string {
	return errsMap[code]
}
