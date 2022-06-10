package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
pool digunakan untuk menyimpan data, kita bisa mengambil data dari pool,
dan setelah selesai kita bisa menyimpan kembali ke pool.
Selama data belum dikembalikan maka data yg diambil akan hilang dr pool.
Istilah lainya seperti data dipinjam dan dikembalikan
*/

func TestPool(t *testing.T) {
	var group sync.WaitGroup
	pool := sync.Pool{}

	// membuat default data pool jika datanya kosong (belum ada yg dikembalikan)
	// pool := sync.Pool{
	// 	New: func() interface{} {
	// 		return "Default data pool"
	// 	},
	// }

	pool.Put("Lukman")
	pool.Put("Sanjaya")
	pool.Put("Ok")

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()

			group.Add(1)
			data := pool.Get()
			fmt.Println(data) // akan ada hasil print yg nil, karena data belum selesai dikembalikan
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	group.Wait()
}
