package go_context

import (
	"context"
	"fmt"
	"testing"
)
/**
Membuat context
- Karena context adalah sebuah interface, untuk membuat context kita butuh sebuah struct yang sesuai dengan kontrak interface Context
- Namun kita tidak perlu membuat secara manual
- Di golang, package context terdapat function yang bisa kita gunakan untuk membuat Context

Function membaut context
- context.Background()		= Membuat context kosong. Tidak pernah dibatalkan, tidak pernah timeout, dan tidak memiliki value apapun.
							  Biasanay digunakan di main function atau dalam test, atau dalam awal proses request terjadi.
- context.TODO()			= Membuat context kosong seperti Background(), namun biasanya penggunaan todo digunakan ketika belum jelas
							  context apa yang ingin digunakan.
 */

/**
Parent dan Child Context
- Context menganut konsep parent dan child
- Artinya saat kita membaut context, kita bisa membuat child context dari context yang sudah ada
- Parent context bisa memiliki banyak child, namun child hanya bisa memiliki1 parent context saja
- Konsep ini miirp seperti dengan pewarisan di pemograman berorientasi object

Contoh parent dan child
- Parent A memiliki child B dan C
- Parent B memiliki child D dan E
- Parent C memiliki child F

Hubungan antara Parent Context dan Child Context
- Parent dan Child akan selalu terhubung
- Misal suatu saat kita akan membatalkan Context A, maka semua child (Context B, C) dan semua sub child (Context D,E,F) dari context A akan dobatalkan juga
- Namun jika misal hanya membatalkan context B, maka hanya child (Context E, D) dan semua sub child dari context B yang dibatalkan. parent context B tidak ikut dibatalkan
- Begitu juga nanti saat kita menyisipkan data ke dalam Context A, semua child dan sub child ntya bisa mendapatkan data tersebut
- Namun jika kita menyisipkan data di context B, maka hanya context B dan semua child dan sub childnya yang mendapatkan data, Parent context B tidak akan mendapatkan datanya

Immutable
- Context merupakan object yang immutable, artinya setelah contect dibuat, dia tidak bisa diubah lagi
- Ketika kita menambahkan value ke dalam context, atau menambahkan pengaturan timeout dan lainnya. secara otomatis akan membentuk child context baru, bukan merubah context tersebut
 */
func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

/**
Context With Value
- Pada saaat awal membuat context, context tidak memiliki value. Background() dan TODO() itu defaultnya kosong
- Kita bisa menambahkan value dengan data pair (key -value) ke dalam context
- Saat kita menambahkan value ke context, secara otomatis akan tercipta child context baru, artinya original contextnya tidak akan berubah sama sekali. karena ini sifatnya immutable
- Untuk menambah value ke context, kita bisa menambahkan function context.WithValue(parent, key, value)
 */


