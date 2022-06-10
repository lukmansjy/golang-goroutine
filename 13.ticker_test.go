package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// ticker adalah representasi kejadian yang berulang sesuai waktu yang ditentukan

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(10 * time.Second)
		ticker.Stop()
	}()

	// selama belum mejalankan ticker.Stop() ticker tidak akan pernah berhenti
	for tick := range ticker.C {
		fmt.Println(tick)
	}

}

// hanya akan mengembalikan channelnya saja, tidak dengan tickernya
func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	// tidak akan pernah berhenti
	for tick := range channel {
		fmt.Println(tick)
	}

}
