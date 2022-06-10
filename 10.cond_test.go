package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
cond atau condition adalah implementasi locking,
cond membutuhkan locker (bisa menggunakan Mutex atau RWMutex) untuk implementasi locking nya.
namun berbeda dengan locking lainya, cond terdapat function Wait() untuk menunggu, apakah perlu menunggu atau tidak.
function Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu lagi, sedangkan
function Broadcase() digunakan untuk memberitahu semua goroutine agar tidak perlu menunggu lagi
*/

var cond = sync.NewCond(&sync.Mutex{})
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	cond.L.Lock()
	cond.Wait() // akan terus menunggu sampai ada cond.Signal() atau cond.Broadcast()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
}
