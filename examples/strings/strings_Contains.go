package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafoo", "foo"))  // true
	fmt.Println(strings.Contains("seafoo", "food")) // false
	// 字符串都包含「空字符串」
	fmt.Println(strings.Contains("seafoo", "")) // true
	fmt.Println(strings.Contains("", ""))       // true
}
