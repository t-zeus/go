package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("Gopher", "go")) //false
	fmt.Println(strings.HasPrefix("Gopher", "Go")) //true
	fmt.Println(strings.HasPrefix("Gopher", ""))   //true

	fmt.Println(strings.HasSuffix("Gopher", "er")) //true
	fmt.Println(strings.HasSuffix("Gopher", "Er")) //false
	fmt.Println(strings.HasSuffix("Gopher", ""))   //true

}
