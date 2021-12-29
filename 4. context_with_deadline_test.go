package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**
Context With Deadline
- Selain menggunakan timeout untuk melakukan cancel secara otomatis, kita juga bisa menggunakan deadline
- Sebenarnya timeout dan deadline tidak ada bedanya, hanya saja deadline mempunyai parameter waktu untuk setting deadlineya
- Pengaturan deadline sedikit berbeda dengan timeout, jika timeout kita beri waktu atau durasi dari sekarang. kalau deadline
 ditentukan kapan waktu akan berhentinya. misal jam 12 siang, jam 18:15
- Untuk membuat context dengan cancel signal secara otomatis menggunakan deadline. kita bisa menggunakan function context.WithDeadline(parent, time)
*/

func AddCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())
	parent := context.Background()
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(parent, deadline)
	defer cancel() // cancel function tetap dipanggil jika selama timeout belum selesai, terdapat ada error
	fmt.Println("Waktu sekarang: ", time.Now())
	fmt.Println("Deadline: ", deadline)

	// mengirim parameter context
	destination := AddCounter(ctx)
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())

	for data := range destination {
		fmt.Println("Cunter: ", data)
	}

	fmt.Println("Selesai: ", time.Now())
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())
}
