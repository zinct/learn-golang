package main

import (
	"fmt"
	"sync"
	"testing"
)

func Increment(counter *int) {
	*counter++
}

func TestOnce(t *testing.T) {
	once := &sync.Once{}
	waitGroup := &sync.WaitGroup{}
	counter := 0

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func() {
			once.Do(func() {
				Increment(&counter)
			})
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	fmt.Println("Counter", counter)
}