## 09 Concurrent Programing
## Kesimpulan

1. Sequential program yaitu program yang dijalankan pada satu waktu, parallel program yaitu beberapa tugas di dalam program yang dijalankan secara bersamaan, concurrent program yaitu memungkinkan beberapa tugas yang dijalankan secara bersamaan, namun belum tentu tugas di eksekusi secara fisik

2. Gourutine adalah sebuah function atau fitur di dalam go yang memungkinkan menjalankan secara indipenden di dalam sebuah program. Channel adalah sebuah cara yang digunakan untuk berkomunikasi antar gourutine. Select adalah sebuah cara untuk memudahkan mengontrol data komunikasi antara satu atau banyak channel.

3. Race contidion adalah kondisi dimana 2 atau lebih threads yang saling bersaing untuk mengakses memory atau variabel. Dalam hal ini cara untuk mengatasi data race yaitu dengan menggunakan WaitGroups, Channel Blocking, dan Mutex.