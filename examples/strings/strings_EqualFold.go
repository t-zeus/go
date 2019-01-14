package main

import (
	"fmt"
	"strings"
)

func main() {
	// func EqualFold(s, t string) bool 不区分大小写比较字符串
	// func ContainsAny(s, chars string) bool 区分大小写

	fmt.Println(strings.EqualFold("go", "Go")) //true
	fmt.Println(strings.EqualFold("", ""))     //true
	fmt.Println(strings.EqualFold("a", "b"))   //false
}
