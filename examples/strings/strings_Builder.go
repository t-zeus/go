package main

import (
	"fmt"
	"strings"
)

func main() {
	/**
	通过 Write 相关方法创建字符串的一种方式
	*/
	var b strings.Builder
	for i := 0; i < 3; i++ {
		fmt.Fprintf(&b, "%d...", i)
	}
	fmt.Println("len=", b.Len())
	fmt.Println("len=", len(b.String()))
	b.Reset()
	b.WriteString("ignition")
	fmt.Println(b.String())
}
