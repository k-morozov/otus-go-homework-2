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

func generate(out chan<- int) {
	defer trace("generate")()
	defer close(out)
	for x := 0; ; x++ {
		if cancelled() {
			return
		}
		out <- x
	}
}

func square(in <-chan int, out chan<- int) {
	defer trace("square")()
	defer close(out)

	for x := range in {
		out <- x
	}
}

func printer(in <-chan int) {
	defer trace("printer")()
	defer close(done)

	counter := 0
	for x := range in {
		counter++
		fmt.Println(x)
		if counter > 10 {
			break
		}
	}
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	numbers := make(chan int)
	output := make(chan int)

	go generate(numbers)
	go square(numbers, output)
	printer(output)
}
