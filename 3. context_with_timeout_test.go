package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
Context With Timeout
- Selain menambahkan value ke context dan juga sinyal cancel, kita juga bisa menambahkan sinyal cancel ke context secara otomatis
dengan menggunakan pengaturan timeout
- Dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel secara manua. cancel akan otomatis dieksekusi
jika waktu timeout sudah terlewati
- Penggunaan context dengan timeout sangat cocok ketika misal kita melakukan query ke database atau http api, namun ingin menentukan batas maksimum timeoutnya
- Untuk membuat context dengan cancel signal secara otomatis menggunakan timeout, kita bisa menggunakan function context.WithTimeout(parent, duration)
*/

func Counter(ctx context.Context) chan int {
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

func TestContextWithTimeout(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)

	defer cancel() // cancel function tetap dipanggil jika selama timeout belum selesai, terdapat ada error

	fmt.Println("Total goroutine: ", runtime.NumGoroutine())
	// mengirim parameter context
	destination := Counter(ctx)
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())

	for data := range destination {
		fmt.Println("Cunter: ", data)
	}

	fmt.Println("Total goroutine: ", runtime.NumGoroutine())
}
