package main

import (
	"fmt"
	"github.com/k-morozov/otus-go-homework-2/foobar"
)

func main() {
	fmt.Println("Hello world!")
	count := foobar.Count("abcdea", 'a')

	fmt.Println(count)
}
