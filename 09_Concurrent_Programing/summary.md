## 09 Concurrent Programing

## Different Sequential Program, Parallel Program, Concurrent Program
* Sequential program yang menjalankan instruksi satu per satu, berurutan, tanpa paralelisme.
* Parallel program yang menjalankan tugas secara bersamaan pada banyak prosesor, untuk meningkatkan kinerja.
* Concurrent program yang menjalankan tugas secara bersamaan, dikelola oleh sistem operasi, tidak selalu melibatkan banyak prosesor. Digunakan untuk mengatasi tugas-tugas bersamaan, termasuk yang melibatkan I/O.

## Concurrency di Go
Concurrency di Go memungkinkan paralelisme. Ini berarti bahwa Anda dapat menjalankan banyak tugas atau goroutine secara bersamaan, yang pada gilirannya memungkinkan paralelisme eksekusi instruksi pada tingkat yang lebih rendah.

## Go Feature
* Concurrent execution (Goroutines).
Goroutine adalah sebuah function atau fitur di dalam go yang memungkinkan menjalankan secara indipenden di dalam sebuah program.

* Synchronization and messaging (Channels).
Channel adalah sebuah cara yang digunakan untuk berkomunikasi antar Goroutine

* Multi-way concurrent control (Select)
Select adalah sebuah cara untuk memudahkan mengontrol data komunikasi antara satu atau banyak channel.

## Contoh Goroutines
```
package main

import (
    "fmt"
    "time"
)

func foo() {
    for i := 0; i < 5; i++ {
        fmt.Println("Foo:", i)
    }
}

func main() {
    go foo()

    time.Sleep(time.Second)
    fmt.Println("Selesai")
}
```

## Contoh Select
```
package main

import "fmt"

func main() {
    // Membuat channel untuk mengirim dan menerima data bertipe int
    ch := make(chan int)

    // Menggunakan select untuk mengirim dan menerima data ke/dari channel
    select {
    case ch <- 42: // Mengirim data 42 ke channel
        fmt.Println("Data berhasil dikirim ke channel.")
    case data := <-ch: // Menerima data dari channel
        fmt.Println("Menerima data dari channel:", data)
    default:
        fmt.Println("Tidak ada aksi yang diambil.")
    }
}
```

## GOMAXPROCS
GOMAXPROCS adalah digunakan untuk mengendalikan jumlah thread sistem operasi yang digunakan untuk menjalankan Goroutine. Fungsi ini memungkinkan untuk menentukan berapa banyak thread sistem operasi yang akan digunakan oleh runtime Go untuk menjalankan konkurensi.

## Unbuffered dan Buffered Channels
* Unbuffered Channel: Tidak memiliki kapasitas buffer. Pengiriman data langsung memblokir hingga ada penerima yang siap. Cocok untuk komunikasi langsung dan sinkron.

* Buffered Channel: Memiliki kapasitas buffer yang dapat Anda tentukan. Pengiriman data tidak langsung memblokir selama buffer belum penuh, membantu mengurangi latensi dan mengizinkan komunikasi asinkron.

```
*Unbuffered
c := make(chan string)

* Buffered
c := make(chan string, 42)
```

## Race Condition
Race contidion adalah kondisi dimana 2 atau lebih threads yang saling bersaing untuk mengakses memory atau variabel. Dalam hal ini cara untuk mengatasi data race yaitu dengan menggunakan WaitGroups, Channel Blocking, dan Mutex.

## WaitGroups
Blocking with WaitGroups adalah teknik yang digunakan dalam bahasa pemrograman Go untuk menunggu sejumlah goroutine selesai sebelum melanjutkan eksekusi program. sync.WaitGroup adalah tipe data yang disediakan oleh paket sync yang digunakan untuk mengoordinasikan penungguan atau penghitungan sejumlah goroutine tertentu.

```
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Menambahkan goroutine ke WaitGroup
        go func(id int) {
            defer wg.Done() // Mengurangkan hitungan WaitGroup saat goroutine selesai
            fmt.Printf("Goroutine %d selesai\n", id)
        }(i)
    }

    // Menunggu semua goroutine selesai
    wg.Wait()
    fmt.Println("Semua goroutine selesai.")
}
```

## Channel Blocking
Channel blocking dalam Go terjadi ketika pengiriman atau penerimaan data melalui channel membuat program berhenti atau terblokir sementara. Channel blocking ini mengizinkan goroutines kita untuk sinkronkan satu sama lain sejenak. 

```
package main

import "fmt"

func main() {
    ch := make(chan int) // Membuat unbuffered channel
    go func() {
        data := 42
        ch <- data // Mengirim data ke channel (akan memblokir jika tidak ada yang menerima)
    }()
    result := <-ch // Menerima data dari channel
    fmt.Println("Hasil:", result)
}
```

## Mutex
Mutex memastikan bahwa hanya satu goroutine pada satu waktu yang dapat mengakses data bersama. Mutex digunakan untuk mengunci (lock) akses ke data bersama saat digunakan dan membuka kunci (unlock) setelah selesai. Ini mencegah race condition dan konflik akses, memastikan data bersama diakses secara aman.

```
package main

import (
    "fmt"
    "sync"
)

var data int
var mutex sync.Mutex // Membuat mutex

func main() {
    // Goroutine 1
    go func() {
        mutex.Lock()   // Mengunci mutex
        data = 42       // Mengakses dan memodifikasi data bersama
        mutex.Unlock() // Membuka kunci mutex
    }()

    // Goroutine 2
    go func() {
        mutex.Lock()   // Mengunci mutex
        result := data // Mengakses data bersama
        mutex.Unlock() // Membuka kunci mutex
        fmt.Println("Hasil:", result)
    }()

    // Menunggu sejenak agar goroutine selesai
    fmt.Println("Menunggu sejenak...")
}
```
