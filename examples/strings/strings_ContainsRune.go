package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsRune("abcdef", 97))  // true
	fmt.Println(strings.ContainsRune("dddd", 97))    // false
	fmt.Println(strings.ContainsRune("Go编程", 32534)) // true  编->32534
}
