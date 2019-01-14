package main

import (
	"fmt"
	"strings"
)

/**
func ContainsAny(s, chars string) bool
字符串中是否包含 字符列表 中的某个字符
*/
func main() {
	fmt.Printf("%d\n", '编')                              // 32534
	fmt.Println(strings.ContainsAny("team", "i"))        // false
	fmt.Println(strings.ContainsAny("failure", "u & i")) // true
	// 测试是否包含 空
	fmt.Println(strings.ContainsAny("team", "")) // false
	// 测试是否 空包含空
	fmt.Println(strings.ContainsAny("", "")) // false
	// 测试包含中文
	fmt.Println(strings.ContainsAny("Go 编程", "编")) // true
}
