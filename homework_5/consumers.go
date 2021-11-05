package homework_5

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type threadError struct {
	m       sync.Mutex
	er      error
	counter int
}

func Run(tasks []Task, countGo, countError int) error {
	workers := make(chan struct{}, countGo)
	defer close(workers)

	var globalError = threadError{}

	var wg sync.WaitGroup

	for _, task := range tasks {
		globalError.m.Lock()
		if globalError.counter >= countError {
			globalError.er = ErrErrorsLimitExceeded
			globalError.m.Unlock()
			break
		}
		globalError.m.Unlock()

		wg.Add(1)
		workers <- struct{}{}

		go runTask(task, &globalError, func() {
			<-workers
			wg.Done()
		})
	}

	fmt.Println("Wait")
	wg.Wait()

	return globalError.er
}

func runTask(task Task, globalError *threadError, handler func()) {
	defer handler()

	err := task()
	if err != nil {
		globalError.m.Lock()
		globalError.counter++
		globalError.m.Unlock()
	}
}
