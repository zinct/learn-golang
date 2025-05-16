package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	counter := 0
	mutex := &sync.Mutex{}

	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter", counter)
}
