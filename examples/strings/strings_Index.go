package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 查不到，返回 -1
	fmt.Println(strings.Index("abcd", "bcd"))  //1
	fmt.Println(strings.Index("abcd", "bcde")) //-1

	// 空子串和第一个字符的位置相同，都是 0
	fmt.Println(strings.Index("abcd", "a")) //0
	fmt.Println(strings.Index("abcd", ""))  //0

	fmt.Println(strings.Index("Go编程", "编程"))

	// IndexAny
	// func IndexAny(s, chars string) int：返回 chars 中的字符，在 s 中最早的那个位置
	// 没有返回 -1

	/**
	a：没有
	e：5
	i：2
	o：没有
	u：没有
	v：没有
	故返回：2
	*/
	fmt.Println(strings.IndexAny("chicken", "aeiouy")) //2
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))   //-1

	// 查询，某个 byte 在字符串的位置
	fmt.Println(strings.IndexByte("golang", 'g'))  //0
	fmt.Println(strings.IndexByte("gophers", 'h')) //3
	fmt.Println(strings.IndexByte("golang", 'x'))  //-1

	// 查询，某个 rune 在字符串的位置
	fmt.Println(strings.IndexRune("chicken", 'k')) //4
	fmt.Println(strings.IndexRune("chicken", 'd')) //-1
	fmt.Println(strings.IndexRune("Go编程", '编'))    //2

	// func IndexFunc(s string, f func(rune) bool) int
	// 通过自定义函数，来查询索引位置
	f := func(r rune) bool {
		return unicode.Is(unicode.Han, r)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))    //7
	fmt.Println(strings.IndexFunc("Hello, World", f)) //-1
}
