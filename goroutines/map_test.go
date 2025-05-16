package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	var data sync.Map
	waitGroup := &sync.WaitGroup{}
	
	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(i int) {
			data.Store(i, i)
			waitGroup.Done()
		}(i)
	}

	waitGroup.Wait()
	fmt.Println("Done")
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}