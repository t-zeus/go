package main

import (
	"fmt"
	"strings"
)

func main() {
	/**
	字符串比较的方式：
		1、string.Ccompare 方式，返回 0，1，-1
		2、内建运算符比较，返回 true、false
	*/
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("ab", "a"))
	fmt.Println("a" > "b")
	fmt.Println("ab" > "a")
}
