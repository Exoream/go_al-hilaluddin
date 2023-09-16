## 15 Intro RESTful API
## Kesimpulan

1. Application Programming Interface (API) adalah kumpulan aturan dan protokol yang memungkinkan berbagai perangkat lunak atau aplikasi untuk berkomunikasi satu sama lain. Cara kerja API yaitu client mengirimkan request ke server dan nantinya server akan mengirimkan respon kembali ke client.

2. API berfungsi untuk memungkinkan aplikasi, sistem atau layanan yang berbeda untuk dapat berkomunikasi dan bekerja satu sama lain, memungkinkan pertukaran data dan fungsionalitas yang terstruktur dan terdokumentasi, dan masih banyak fungsi lainnya.

3. Representational State Transfer (REST) adalah arsitektur berbasis standar yang digunakan dalam pengembangan web untuk merancang sistem yang ringan, scalable, dan mudah dipahami. Ini berarti bahwa REST memberikan pedoman tentang cara merancang API yang ringan, skalabel, dan mudah dipahami.

4. Dalam pengembangan API, REST biasanya digunakan bersama dengan protokol HTTP (Hypertext Transfer Protocol) untuk berkomunikasi. REST menggunakan metode HTTP seperti GET, POST, PUT, dan DELETE untuk mengakses dan mengelola sumber daya yang diidentifikasi oleh URL. Selain itu menggunakan format JSON, XML, dan SOAP, namun secara umum kita menggunakan JSON

5. HTTP memiliki sebuah respon code yaitu jika angkanya depannya 200 artinya berhasil, jika angka depannya 400 artinya gagal dari sisi client dan jika angka depannya 500 artinya gagal dari segi server.

6. API memiliki beberapa tools yaitu Katalon, JMeter, SoapUI, Postman, Insomnia, dan masih banyak lainnya.

7. Dalam mendesaign REST API ada beberapa hal yang perlu diperhatikan yaitu gunakan kata benda dari pada kata kerja, menggunakan kata benda jamak, menggunakan pengumpulan sumber daya (resource collection), lebih baik gunakan format JSON, melakukan filtering, sorting, paging, dan field selection, menangani garis miring dengan hati - hati, memungkinkan perubahan dan perkembangan API tanpa mengganggu pengguna yang ada, dan selalu membuat API dokumentasi

8. Package Go "net/http" bertujuan menyediakan implementasi klien dan server http. Package Go "encoding/json" memungkinkan untuk mengubah data Go yang kompleks menjadi format JSON dan sebaliknya. Anda dapat menggunakan struktur data dalam Go untuk merepresentasikan objek JSON, dan paket ini akan mengurus konversi antara dua format tersebut.