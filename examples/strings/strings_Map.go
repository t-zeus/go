package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//func Map(mapping func(rune) rune, s string) string
	//返回经过函数处理后的字符串
	//HELLO,WORLD
	fmt.Println(strings.Map(func(r rune) rune {
		return unicode.ToUpper(r)
	}, "hello,world"))
}
