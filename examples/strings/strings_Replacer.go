package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	s := r.Replace("This is <b>HTML</b>!")
	fmt.Println(s) // This is &lt;b&gt;HTML&lt;/b&gt;!
}
