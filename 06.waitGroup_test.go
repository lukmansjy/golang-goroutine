package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Wait Group digunakan untuk menunggu semua goroutine selesai prosesnya

func RunAsynchronous(group *sync.WaitGroup, i int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello", i)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group, i)
	}

	group.Wait()
	fmt.Println("Selesai")
}
