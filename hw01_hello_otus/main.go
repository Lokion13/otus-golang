package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	inputStr := "Hello, OTUS!"
	reversedStr := reverse.String(inputStr)
	fmt.Println(reversedStr)
}
