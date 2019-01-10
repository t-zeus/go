package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	/**
	如果 Reader 的长度没有 min 那么大，读完之后会返回 ErrUnexpectedEOF.
	*/
	reader := strings.NewReader("123456789")
	buf := make([]byte, 30)
	n, err := io.ReadAtLeast(reader, buf, 30)
	fmt.Println(n, err) // 9 unexpected EOF
	fmt.Printf("%q\n", buf)
}
