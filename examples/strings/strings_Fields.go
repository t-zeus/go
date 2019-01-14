package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//func Fields(s string) []string
	//通过空白分来分割字符串
	fmt.Printf("Fields are: %q\n", strings.Fields("   a  b   \t c")) //Fields are: ["a" "b" "c"]

	//func FieldsFunc(s string, f func(rune) bool) []string
	//通过指定函数来分割字符串
	fmt.Printf("Fields, are: %q\n", strings.FieldsFunc("sfsdf;boo2;farrsdm_$;", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) //Fields, are: ["sfsdf" "boo2" "farrsdm"]
	}))
}
