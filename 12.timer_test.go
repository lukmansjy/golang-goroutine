package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// timer seperti delay job

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time) // perbedaan waktu 5 detik
}

func TestTimeAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time) // perbedaan waktu 5 detik
}

// menjalankan function dengan delay waktu tertentu
func TestTimeAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(3*time.Second, func() {
		defer group.Done()
		fmt.Println("Execute after 3 second")
	})

	group.Wait()
}
