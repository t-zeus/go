package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	var (
		n   int
		err error
	)

	r := strings.NewReader("some")
	buf := make([]byte, 4)
	if n, err = io.ReadFull(r, buf); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("n=%d,buf=%q\n", n, buf)

	buf2 := make([]byte, 4)
	if n, err = io.ReadFull(r, buf2); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("n=%d,buf=%q\n", n, buf2)

	// minimal read size bigger than io.Reader stream
	r2 := strings.NewReader("some2")
	buf3 := make([]byte, 8)
	if n, err = io.ReadFull(r2, buf3); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("n=%d,buf=%q\n", n, buf3)
}
