package main

import "fmt"

func main() {
	// strings 底层是 byte 数组存放，这里面中文占用 3 个字节，故 len(s) = 8
	s := "Go编程"
	fmt.Println(len(s))

	// 如果想要几个字符的话，需要把 string 转换成 rune 切片，再计算长度
	// 此时长度为 4
	runeSlice := []rune(s)
	fmt.Println(len(runeSlice))

	// 计算中文字符，所占用的字节数
	// len：支持的类型有 Array Slice String Pointer Channel
	// 所以，需要将 中文 rune 转换成 string 再计算长度
	fmt.Println(len(string('编'))) // 3
}
