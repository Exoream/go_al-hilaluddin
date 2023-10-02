# 20 Clean and Hexagonal Architecture

## Clean Architecture
Clean Architecture adalah sebuah pendekatan untuk merancang sistem perangkat lunak yang bertujuan untuk menjaga kejelasan, pemisahan tanggung jawab, dan skalabilitas dalam pengembangan perangkat lunak. Misi Clean Architecture adalah untuk menciptakan sistem perangkat lunak yang dapat diandalkan, mudah dipahami, dan mudah diubah dengan memisahkan elemen-elemen inti dari sistem dari detail implementasi. Hal ini bertujuan untuk mempromosikan ketertiban, kejelasan, dan perawatan yang berkelanjutan dalam pengembangan perangkat lunak dengan memprioritaskan pemisahan konsep, tanggung jawab yang jelas, dan pengujian yang kuat. Clean Architecture membantu kita menciptakan sistem yang bisa beradaptasi dengan perubahan kebutuhan bisnis dan teknologi tanpa mengorbankan stabilitas dan kualitas.

## MVC dan MVVM
MVC (Model-View-Controller):
* Model: Menangani data dan logika bisnis.
* View: Bertanggung jawab untuk tampilan dan interaksi dengan pengguna.
* Controller: Mengatur interaksi antara Model dan View serta alur aplikasi.

MVVM (Model-View-ViewModel):
* Model: Menangani data dan logika bisnis seperti dalam MVC.
* View: Menampilkan data dan interaksi dengan pengguna.
* ViewModel: Memisahkan logika tampilan dari View, memungkinkan pengikatan data yang kuat dan pengujian yang lebih mudah.

## Architecture of System
* Hexagonal Arsitektur
* Clean Architecture
* Onion Architecture
* Screaming Architecture
* DCI from Agile Development
* BCE from Object Oriented Software

## Tujuan Mengimplementasikan Clean Architecture
* Independent of Frameworks
* Testable
* Independent of UI
* Independent of Database
* Independent of any external

## CA Layer Main
#### Use Case - Domain layer
* Berisi logika bisnis inti aplikasi.
* Mewakili fungsionalitas khusus yang dapat dilakukan dalam aplikasi dan menjalankan operasi-operasi terkait bisnis.

#### Controller - Presentation layer
* Berfungsi sebagai penghubung antara antarmuka pengguna (UI) dan lapisan bisnis inti.
* Menerima input dari pengguna atau sistem eksternal, menginterpretasikan input tersebut, dan mengarahkan alur aplikasi.

#### Repository
* Menghubungkan lapisan bisnis inti dengan penyimpanan data seperti basis data atau penyimpanan berkas.
* Menyediakan abstraksi untuk operasi data dan mengisolasi lapisan bisnis dari detail teknis penyimpanan data.

## Domain Driven Design (DD)
Domain Driven Design adalah pendekatan dalam pengembangan perangkat lunak yang berfokus pada pemahaman yang mendalam tentang domain bisnis atau domain aplikasi yang sedang dibangun. Tujuan utama DDD adalah menciptakan perangkat lunak yang mencerminkan dan melayani kebutuhan bisnis dengan baik.

## Context Golang
Context adalah paket atau pustaka bawaan (standard library) yang digunakan untuk mengelola nilai-nilai konteks yang berkaitan dengan suatu permintaan atau operasi. Nilai-nilai konteks ini dapat digunakan untuk mengendalikan dan melacak bagaimana suatu operasi atau permintaan berperilaku dalam lingkup yang lebih besar, seperti pengendalian waktu, pembatalan, atau pengiriman informasi yang relevan ke berbagai goroutine dalam aplikasi Go.

## Implementasi Context
```
var ctx = context.Background()
ctx = context.WithValue(ctx, "key", "value")
```

## Tujuan Penggunaan Clean Code
* Modular artinya bisa dengan mudah mengganti dipendensi satu ke yang lain
* Scallabel artinya kita bisa dengan mudah menambah feature baru dan lain sebagainya
* Maintainable artinya kita dengan mudah memperbaiki masalah pada kode kita

## 3 Option Migrasi dari MVC ke Clean Code
* Pertahankan desaign dengan memisahkan depedensi
* Pertahankan desaign tetapi memindahkan ke dalam suatu layer
* Ubah desaign dan pisahkan depedensi