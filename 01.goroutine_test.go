package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

// jalankan dengan unit test, perintah: go test -v -run=TestCreateGoroutine
func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // running menggunakan goroutine (di js disebut asynchronous)
	fmt.Println("Ops")

	time.Sleep(1 * time.Second)
}
