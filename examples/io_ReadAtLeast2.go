package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

/**
官方 demo 稍作修改
link: http://localhost:8000/pkg/io/#ReadAtLeast
*/
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 33)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// buffer smaller than minimal read size.
	// len(shortBuf) < min 这个直接错误了
	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
		fmt.Println("error:", err) // error: short buffer
	}

	// minimal read size bigger than io.Reader stream
	// r 这个 read 已经读取完了，再次调用 r.Read 实际上会直接返回 EOF
	longBuf := make([]byte, 64)
	if n, err := io.ReadAtLeast(r, longBuf, 64); err != nil {
		fmt.Println("error:", err, "\tn:", n) //error: EOF 	n: 0
	}
	// 如果 read size > io.Reader
	r2 := strings.NewReader("some io.Reader stream to be read\n")
	longBuf2 := make([]byte, 64)
	if n, err := io.ReadAtLeast(r2, longBuf2, 64); err != nil {
		fmt.Println("error:", err, "\tn:", n) //error: unexpected EOF 	n: 33
	}
}
