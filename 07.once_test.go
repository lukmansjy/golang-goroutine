package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// once diginakan untuk memastikan sebuah function dijalankan sekali oleh goroutine yg pertama
var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			once.Do(OnlyOnce) // function jalan hanya sekali (hanya bisa func yg tdk memiliki parameter)
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter) // hasilnya 1
}
