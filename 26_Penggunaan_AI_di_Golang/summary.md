# Penggunaan AI di Golang

## Alasan Memilih Go
* Dikembangkan oleh Google untuk pengembangan aplikasi server-side dan berbasis cloud
* Go adalah compiled language sehingga lebih cepat dibanding interpreted language (contoh : Python)
* Mendukung concurrency melalui fitur goroutines sehingga dapat menjalankan banyak task secara simultan tanpa membebani memori
* Memiliki standard libraries yang sangat kaya, didukung oleh komunitas yang kuat
* Termasuk libraries untuk mendukung pengembangan AI

## Library Go Untuk Machine Learning
### GoLearn
* Bersifat open source
* 'Batteries included' machine learning library
* Simplicity & customizability
* Contoh Penerapan : 
[link](https://github.com/sjwhitworth/golearn#getting-started)

### Gorgonia
* Memudahkan menulis dan mengevaluasi persamaan matematika yang melibatkan array multidimensi
* Low-level library, seperti Theano, namun lebih tinggi dibanding Tensorflow
* Contoh Penerapan : 
[link](https://github.com/gorgonia/gorgonia#usage)

### Gonum
* Sebuah set packages yang didesain untuk menulis numerical & algoritma sains menjadi produktif, memiliki performa tinggi, dan scalable
* Mirip seperti numpy dan spicy, library yang dibuat menggunakan Python
* Contoh Penerapan :
[link](https://www.gonum.org/post/intro_to_gonum/)

### Gomind
* Contoh Penerapan :
[link](https://github.com/surenderthakran/gomind#usage)


## AI as a services (AIaas)
## Definisi
* Sebuah tools AI yang dapat segera digunakan (pre-built or off-the-shelf AI tools)
* Berguna bagi bisnis yang ingin menerapkan AI tanpa merekrut ahlinya dan tanpa mengeluarkan biaya yang relatif banyak

Beberapa perusahaan yang menyediakan AIaas, contohnya :
* Amazon Web Services (AWS)
* Microsft Azure
* Google Cloud Platform (GCP)
* IBM
* OpenAI

## Tipe-tipe AI as a services (AIaas)
### Bots/Chatbots : Membuat AI berbasis costumer service
* Amazon lex (AWS), Azure Bot Service (Microsft Azure), DialogFlow(GCP)

### APIs : mengintegrasikan AI milik vendor dengan aplikasi sendiri
* Amazon Rekognition, Amazon Comprehend
* Azure Cognitive Service, Azure Speech Services
* Google Cloud Vision, Cloud Natural language Machine learning

### Machine Learning
* Amazon SageMaker
* Azure Machine learning
* Google Cloud AI

## Keuntungan AI as a services (AIaas)
* Pengurangan Cost
* Low-mode
* Proses deployment cepat
* Flexibility
* Usability
* Scalability
* Customization

## Kelemahan AI as a services (AIaas)
* Cost yang berlebih dalam jangka panjang
* Privasi dapat
* Keamanan
* Transparasi
* Vendor lock-in

## OpenAI API
* Mendesain prompt yang tepat (prompt engineering)
* Membuat aplikasi berbasis AI menggunakan API OpenAI
* [Dokumentasi OpenAI API](https://platform.openai.com/docs/introduction)
* [OpenAI API Reference](https://platform.openai.com/docs/api-reference/introduction)
* Opsional : install library [Go OpenAI](https://github.com/sashabaranov/go-openai)

## OpenAI API Pricing
* Start for free , mendapatkan free credit senilai $5 yang dapat digunakan selama 3 bulan pertama semenjak register
* Pay as you go, setiap model memiliki tarif yang berbeda berbeda
* Lihat detail di [link](https://openai.com/pricing)
* Cek seberapa banyak penggunaan di [link](https://platform.openai.com/account/usage)
* Best practices manajemen billing : [link](https://platform.openai.com/docs/guides/production-best-practices/managing-billing-limits)

## Cara Mendapatkan API Key Untuk Akses OpenAI API
* Buka situs web OpenAI [link](https://platform.openai.com/docs/api-reference)
* Lakukan login, jika belum punya akun silahkan register terlebih dahulu
* Cari bagian introduce dan masuk ke API keys atau masuk ke halaman berikut [link](https://platform.openai.com/account/api-keys)
* Setelah masuk kehalaman tersebut, pilih API keys dan Create new secret key
* Ikuti petunjuk untuk mengkonfigurasi API keys
* Jangan lupa salin secret keys nya, secret keys hanya bisa tampil satu kali

## Promp Engineering : Panduan Penulisan Prompt
### Prinsip 1 : Gunakan Prompt Yang Jelas dan Spesifik
* Gunakan delimiters untuk menandakan bagian mana yang menjadi input
```
(contoh: ```, ''', <>, <tag></tag>)
```
* Tuliskan struktur output yang diinginkan
* Minta model untuk memeriksa apakah input sudah sesuai
* Few-shot prompting

### Prinsip 2 : Berikan "waktu berpikir" Untuk Menghindari Solusi Yang Salah
* Menulis langkah-langkah yang dibutuhkan untuk menyelesaikan tugas dan output yang diinginkan
* Menginstruksikan model untuk menuliskan solusinya sendiri sebelum masuk ke kesimpulan
 