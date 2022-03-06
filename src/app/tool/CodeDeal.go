package tool

import (
	"crypto/md5"
	"encoding/hex"
)

//密码处理

func Encryption(code string) string {
	passwd := []byte(code)
	passwd = append(passwd, passwd[0], passwd[2])
	ta := md5.New()
	ta.Write(passwd)
	return hex.EncodeToString(ta.Sum(nil))
}

//密码对比  code01是输入的密码 code02是database里的密码
func CompasionCode(code01 string, code02 string) bool {
	passwd := []byte(code01)
	passwd = append(passwd, passwd[0], passwd[2])
	ta := md5.New()
	ta.Write(passwd)
	code01 = hex.EncodeToString(ta.Sum(nil))
	if code01 != code02 {
		return false
	}
	return true
}
