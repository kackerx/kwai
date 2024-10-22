package code

import "fmt"

type Merror struct {
	code int
	msg  string
}

var errMap map[int]Merror

func registerErrCode(code int, msg string) {
	errMap[code] = Merror{code, msg}
}

func GetErrResp(code int) Merror {
	return errMap[code]
}

func GetErrMsg(code int, args ...any) string {
	return fmt.Sprintf(errMap[code].msg, args...)
}
