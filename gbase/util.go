package gbase

import "strings"

type baseUtil struct{}

var Util = &baseUtil{}

func (baseUtil) Add(v1 int, v2 int) int {
	return v1 + v2
}

func (baseUtil) AddStr(v1 string, v2 string) string {
	return v1 + v2
}

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
