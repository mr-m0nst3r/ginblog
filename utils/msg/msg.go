package msg

const (
	SUCCESS = 200
	ERROR = 500

	// code = 1000... user msg

	ErrUsernameUsed   = 1001
	ErrPasswordWrong  = 1002
	ErrUserNotExist   = 1003
	ErrTokenNotExist  = 1004
	ErrTokenExpired   = 1005
	ErrTokenWrong     = 1006
	ErrTokenTypeWrong = 1007

	// code = 2000... article msg


	// code =  3000... category msg
	ErrCategoryUsed = 3001
)

var Codemsg = map[int]string{
	SUCCESS:           "OK",
	ERROR:             "FAIL",
	ErrUsernameUsed:   "用户名已存在",
	ErrPasswordWrong:  "密码错误",
	ErrUserNotExist:   "用户不存在",
	ErrTokenNotExist:  "TOKEN不存在",
	ErrTokenExpired:   "TOKEN已过期",
	ErrTokenWrong:     "TOKEN不正确",
	ErrTokenTypeWrong: "TOKEN格式错误",
	ErrCategoryUsed: "分类名已被占用",
}

func GetMsg(code int) string  {
	return Codemsg[code]
}
