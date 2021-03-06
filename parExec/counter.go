package parExec

import "sync"

type threadSafeCounter interface {
	add()
	check() bool
}

type threadSafeCounterImpl struct {
	m              sync.Mutex
	counter        int
	maxCountErrors int
}

func (c *threadSafeCounterImpl) add() {
	c.m.Lock()
	defer c.m.Unlock()

	c.counter++
}

func (c *threadSafeCounterImpl) check() bool {
	c.m.Lock()
	defer c.m.Unlock()
	if c.counter >= c.maxCountErrors {
		return false
	}
	return true
}

func newThreadSafeCounter(maxErrors int) threadSafeCounter {
	return &threadSafeCounterImpl{maxCountErrors: maxErrors}
}
