package mylib

import (
	"fmt"
	"strings"
)

func init() {
	// 每个源文件都可以定义init函数, 导入包时调用
	fmt.Println("init")
}

func Partition(s string, sep string) (head string, retSep string, tail string) {
	// Partition(s, sep) -> (head, sep, tail)
	index := strings.Index(s, sep)
	if index == -1 {
		head = s
		retSep = ""
		tail = ""
	} else {
		head = s[:index]
		retSep = sep
		tail = s[len(head)+len(sep):]
	}
	return
}
