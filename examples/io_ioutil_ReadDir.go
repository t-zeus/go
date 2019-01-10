package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalln(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Print("[DIR] ")
		} else {
			fmt.Print("[FILE] ")
		}
		fmt.Println(file.Name())
	}
}
