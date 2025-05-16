package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var counter int64 = 0
	waitGroup := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Println("Counter", counter)
}