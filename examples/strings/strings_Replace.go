package main

import (
	"fmt"
	"strings"
)

func main() {
	//func Replace(s, old, new string, n int) string
	//n替换次数，如果是-1，替换所有
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      //oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //moo moo moo
}
