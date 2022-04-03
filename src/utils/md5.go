package utils

import (
	"crypto/md5"
	"fmt"
)

func PwdHash(pwd string) (pwdHash string) {
	data := []byte(pwd)
	pwdHash = fmt.Sprintf("%x", md5.Sum(data))
	fmt.Println(pwdHash)
	return pwdHash
}

func CheckPwd(pwd string, pwdHash string) (res bool) {
	data := []byte(pwd)
	pwdRes := fmt.Sprintf("%x", md5.Sum(data))
	if pwdRes != pwdHash {
		return false
	}
	return true
}
