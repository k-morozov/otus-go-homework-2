package homework_5

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, maxCountGo, maxCountErrors int) (errorResult error) {
	workers := make(chan struct{}, maxCountGo)
	defer close(workers)

	counterErrors := newThreadSafeCounter(maxCountErrors)
	var wg sync.WaitGroup

	for _, task := range tasks {
		if !counterErrors.check() {
			errorResult = ErrErrorsLimitExceeded
			break
		}

		wg.Add(1)
		workers <- struct{}{}

		go runTask(task, counterErrors, func() {
			<-workers
			wg.Done()
		})
	}

	wg.Wait()

	return errorResult
}

func runTask(task Task, counterErrors threadSafeCounter, handler func()) {
	defer handler()

	err := task()
	if err != nil {
		counterErrors.add()
	}
}
