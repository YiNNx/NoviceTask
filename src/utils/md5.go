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
