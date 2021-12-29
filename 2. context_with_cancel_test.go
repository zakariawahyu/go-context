package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**
Context With Cancel
- Selain menambah value ke context, kita juga bisa mengirim sinyal cancel ke context
- Biasanya ketika kita butuh menjalankan proses lain dan kita ingin memberi sinyal cancel ke proses tersebut
- Biasanya proses ini berupa goroutine berbeda, sehingga dengan mudah jika kita ingin membatalkan eksekusi goroutine, kita bisa mengirim sinyal ke contextnya
jika dalam goroutine mendeteksi context cancel, maka akan langsung membatalkan proses selanjutnya
- Namum ingat, goroutine yang menggunakan context tetap harus melakukan pengeceka terhadap contextnya, jika tidak, tidak ada gunanya
- Untuk membuat context dengan cancel signal, kita bisa menggunakan function context.WithCancel(parent)
*/

func CreateCounterLeak() chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}

func TestCounterLeak(t *testing.T) {
	//Contoh counter leak
	fmt.Println("Total goroutine leak: ", runtime.NumGoroutine())
	for data := range CreateCounterLeak() {
		fmt.Println("Cunter: ", data)
		if data == 25 {
			break
		}
	}
	fmt.Println("Total goroutine leak: ", runtime.NumGoroutine()) // walaupun sudah di break go routine masih berjalan, jadi proses ini disebut goroutine leak (terlanjt jalan tidak bisa di kill)
}

func CreateCounter(ctx context.Context) chan int {
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
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	fmt.Println("Total goroutine: ", runtime.NumGoroutine())
	// mengirim parameter context
	destination := CreateCounter(ctx)
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())

	for data := range destination {
		fmt.Println("Cunter: ", data)
		if data == 10 {
			break
		}
	}

	// mengirim signal cancel ke context
	cancel() // jika dipanggil maka context akan dibatalkan

	time.Sleep(2 * time.Second)
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())
}
