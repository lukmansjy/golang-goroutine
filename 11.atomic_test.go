package golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/*
atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent.
contohnya pada sebelumnya kita menggunakan mutex untuk melakukan locking ketika menaingak angka di counter,
hal ini kita bisa menggunakan atomic
*/

func TestAtomic(t *testing.T) {
	group := sync.WaitGroup{}
	var counter int64 = 0
	for i := 1; i <= 1000; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	group.Wait()
	fmt.Println("Counter :", counter)
}
