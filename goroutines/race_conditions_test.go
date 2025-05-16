package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func IncrementCounter(counter *int, mutex *sync.Mutex) {
	for j := 1; j <= 100; j++ {
		mutex.Lock()
		*counter++;
		mutex.Unlock()
	}
}

func TestRaceCondition(t *testing.T) {
	counter := 0;
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go IncrementCounter(&counter, &mutex)
	}

	time.Sleep(5 * time.Second)

	fmt.Println(counter);
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()

	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	// Menggunakan jumlah goroutine yang lebih sedikit
	for i := 1; i <= 10; i++ {
		go func() {
			account.AddBalance(1)
			fmt.Println(account.GetBalance())
		}()
	}

	time.Sleep(1 * time.Second)

	fmt.Println("Final Balance:", account.GetBalance())
}