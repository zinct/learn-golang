package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			data := "Ini data baru"
			return &data
		},
	}
	waitGroup := &sync.WaitGroup{}

	indra := "Indra"
	mahesa := "Mahesa"
	audri := "Audri"
	mona := "Mona"

	pool.Put(&indra)
	pool.Put(&mahesa)
	pool.Put(&audri)
	pool.Put(&mona)
	
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			data := pool.Get()
			fmt.Println(*(data.(*string)))
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	fmt.Println("Done")
}
	