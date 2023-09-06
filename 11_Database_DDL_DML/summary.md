## 11 Database – DDL – DML
## Kesimpulan

1. Database merupakan sebuah kumpulan data yang tersimpan secara terstruktur dan dapat dikelola, diakses, dan diperbaharui melalui query yang tersedia atau secara langsung menggunakan software. Database memiliki relationship yaitu One to One, One to Many, dan Many to Many

2. Secara umum perintah di SQL terdiri atas 3 yaitu Data Definition Language(DDL) yang artinya sebuah perintah yang membangun struktur atau skeleton dari database dan table contohnya CREATE DATABASE/TABLE, USE, DROP, RENAME, ALTER, ADD. 
Data Manipulation Language(DML) yang artinya sebuah perintah untuk memanipulasi data didalam sebuah table atau yang biasa disebut CRUD. Contoh dari DLL yaitu INSERT, SELECT, DELETE, UPDATE selain itu DML memiliki beberapa statement yaitu LIKE/BETWEEN, AND/OR, ORDER BY, LIMIT. Terakhir Data Control Language(DCL) artinya sebuah perintah yang digunakan untuk mengontrol hak akses, izin, dan keamanan dalam basis data. Contohnya GRANT dan REVOKE.

3. Dari ketiga hal diatas sebenarnya ada juga yang dinamakan Transaction Control Language(TCL) yang artinya sebuah perintah untuk mengelola transaksi dalam database. Ini digunakan untuk mengelola perubahan yang dibuat oleh pernyataan DML. Contohnya COMMIT, ROLLBACK, dan SAVEPOINT. 