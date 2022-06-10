package golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {

		group.Add(1)
		go func() {
			defer group.Done()
			time.Sleep(1 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("CPU:", totalCpu)

	totalTread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread:", totalTread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine:", totalGoroutine)

	group.Wait()
}
