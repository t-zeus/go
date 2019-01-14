package main

import (
	"fmt"
	"strings"
)

func main() {
	// 计算指定子串的个数
	fmt.Println(strings.Count("cheese", "e")) // 3 统计包含 e 的个数
	// 1 c 2 h 3 e 4 e 5 s 6 e 7
	fmt.Println(strings.Count("cheese", "")) // 7 统计所有的空，before & after each rune
	fmt.Println(strings.Count("Go编编程", "编")) // 2
	// 1 编 2 程 3
	fmt.Println(strings.Count("编程", "")) // 3
}
