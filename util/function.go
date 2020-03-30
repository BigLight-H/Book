package util

import "strconv"

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

//字符串转换int
func StrToInt(str string) int {
	i,_ := strconv.Atoi(str)
	return i
}
