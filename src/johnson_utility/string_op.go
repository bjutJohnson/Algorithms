package johnson_utility

import (
	"bytes"
)

// 字符串连接操作
func ConcateString(first string, lefts ...string) string {
	var buffer bytes.Buffer

	buffer.WriteString(first)
	for _, v := range lefts {
		buffer.WriteString(v)
	}

	return buffer.String()
}
