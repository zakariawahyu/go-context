**Pengenalan Golang Context**
- Context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout, dan sinyal deadline
- Context biasanya dibuat per request (misal setiap ada request masuk ke server web melalui http request)
- Context digunakan untuk mempermudah kita meneruskan value dan sinyal ke dalam proses

**Kenapa Context Perlu Dipelajari ?**
- Context di Golang biasa digunakan untuk mengirim data request atau sinyal ke proses lain
- Dengan menggunakan Context, ketika kita ingin membatalkan semua proses, kita cukup mengirim sinyal ke Context, maka otoamtis semua proses akan dibatalkan
- Hampir semua bagian di Golang memanfaatkan Context, seperti database, http server, http client, dll
- Bahkan di Google sendiri, ketika menggunakan Golang, context wajib digunakan dan selalu dikirim ke setiap function yang dikirim

**Cara Kerja Context**
- Misal dari proses A mengirim ke proses B dan C
- Didalam proses A terdapat cotext yang akan mengrim sinyal ke proses B dan C
- Jika proses A mengirim sinyal cancel, maka proses B dan C akan menerima sinyal tersebut dan akan membatalkan prosesnya
- Jadi context itu hanya sebatas data saja

**Package Context**
- Context direpresentasikan di dalam sebuah interface Context
- Interface Context terdapat dalam package Contect
- https://golang.org/pkg/context

