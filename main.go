package main

import (
	"fmt"
	"log"
	"time"
)

func trace(text string) func() {
	start := time.Now()
	log.Printf("%s start: %s", text, start)
	return func() {
		log.Printf("%s finish", text)
	}
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	fmt.Println("Hello world!")
}

func main() {
	bigSlowOperation()
}
