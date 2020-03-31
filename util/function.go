package util

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strconv"
)

//返回指定格式的数据
func ReturnJson(msg string, code int, src []map[string]string) map[string]interface{} {
	data := map[string]interface{}{}
	data["data"] = src
	data["message"] = msg
	data["code"] = code
	return data
}

//返回搜索数据map
func BackInfoMap(info []map[string]string, link string, name string, uname string, id string) []map[string]string {
	info = append(
		info,
		map[string]string{"link": link, "name": name, "uname": uname, "id": id})
	return info
}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str) )
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//email verify
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//字符串转换int
func StrToInt(str string) int {
	i,_ := strconv.Atoi(str)
	return i
}
