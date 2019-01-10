package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello,world!")
	// stops with EOF after n bytes
	lr := io.LimitReader(r, 5)
	buf := make([]byte, 8)
	n, err := lr.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("n=%d,buf=%v\n", n, buf)

}
