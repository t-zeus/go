package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "hello,world"
	r := strings.NewReader(s)
	fmt.Println("size=", r.Size())
	buf := make([]byte, 5)
	var (
		n   int
		err error
	)
	n, err = r.Read(buf)
	fmt.Println("n:", n, "err:", err)
	fmt.Printf("buf=%q\n", buf)
	fmt.Printf("size=%d\n", r.Size())
	fmt.Printf("len=%d\n", r.Len())
	r.Reset(s)
	fmt.Printf("len=%d\n", r.Len())

	n, err = r.ReadAt(buf, 6)
	fmt.Printf("buf=%q\n", buf)
	fmt.Printf("len=%d\n", r.Len())

	r.Read(buf)

	r.WriteTo(os.Stdout)
}
