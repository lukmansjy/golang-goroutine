package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", x) // hasil pasti 100rb
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
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

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(userSender *UserBalance, userReceiver *UserBalance, amount int) {
	userSender.Lock()
	fmt.Println("Lock user sender", userSender.Name)
	userSender.Change(-amount)

	time.Sleep(1 * time.Second)

	userReceiver.Lock()
	fmt.Println("Lock user receiver", userReceiver.Name)
	userReceiver.Change(amount)

	userSender.Unlock()
	userReceiver.Unlock()
}

func TestDeadLock(t *testing.T) {
	lukman := UserBalance{
		Name:    "Lukman",
		Balance: 100000,
	}

	budi := UserBalance{
		Name:    "Budi",
		Balance: 100000,
	}

	go Transfer(&lukman, &budi, 50000)
	go Transfer(&budi, &lukman, 60000)

	time.Sleep(3 * time.Second)

	fmt.Println("user", lukman.Name, "Balance", lukman.Balance)
	fmt.Println("user", budi.Name, "Balance", budi.Balance)
}
