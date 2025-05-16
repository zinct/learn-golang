package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Account struct {
	RWMutex sync.Mutex
	name string
	Balance int
}

func (account *Account) Lock() {
	account.RWMutex.Lock()
}

func (account *Account) Unlock() {
	account.RWMutex.Unlock()
}

func (account *Account) ChangeBalance(amount int) {
	account.Balance = account.Balance + amount
}

func Transfer(account1 *Account, account2 *Account, amount int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	account1.Lock()
	fmt.Println("Lock account1", account1.name)
	account1.ChangeBalance(-amount)
	account1.Unlock()

	time.Sleep(1 * time.Second)
	
	account2.Lock()
	fmt.Println("Lock account2", account2.name)
	account2.ChangeBalance(amount)
	account2.Unlock()

	time.Sleep(1 * time.Second)
	fmt.Println("Success transfer", account1.name, "to", account2.name)
}

func TestDeadlock(t *testing.T) {
	account1 := Account{
		name: "Indra Mahesa",
		Balance: 1000,
	}
	account2 := Account{
		name: "Audri Mona Najogi",
		Balance: 1000,
	}

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)

	go Transfer(&account1, &account2, 1000, waitGroup)
	go Transfer(&account2, &account1, 1000, waitGroup)
	
	waitGroup.Wait()
	fmt.Println("Balance account1", account1.Balance)
	fmt.Println("Balance account2", account2.Balance)
}