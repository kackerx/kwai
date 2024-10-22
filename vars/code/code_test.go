package code

import (
	"fmt"
	"testing"
)

func TestGetErrMsg(t *testing.T) {
	fmt.Println(GetErrMsg(ErrUserNotFound, 1, "22"))
	fmt.Println(GetErrMsg(ErrUserAuthFaild))
}
