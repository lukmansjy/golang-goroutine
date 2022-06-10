package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// Map sync ini mirip seperti map golang, yang membedakan adalah map sync ini aman digunakan do concurrent

func addToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestSyncMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go addToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})

}
