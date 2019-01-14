package main

import (
	"fmt"
	"strings"
)

func main() {
	//slice 连接成 string
	//func Join(a []string, sep string) string
	s := []string{"abc", "def", "ghi"}
	fmt.Println(strings.Join(s, "  ")) //abc  def  ghi

}
