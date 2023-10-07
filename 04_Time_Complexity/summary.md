## Time Complexity

## Apa Itu Time Complexity
Time complexity merupakan sebuah konsep algoritma yang bertujuan untuk mengetahui waktu yang diperlukan untuk menjalankan sebuah program atau perintah sesuai dengan inputan yang diberikan. Penggunaan time complexity mengukur agar sebuah program berjalan efisien. Time complexity dijabarkan menggunakan notasi O(Big O).

## Perbandingan Time Complexity
* Constant time - O(1)
Algoritma ini memiliki waktu eksekusi yang konstan, tidak peduli seberapa besar ukuran inputnya. Misalnya, mengakses elemen dalam array menggunakan indeks memiliki time complexity O(1)

* Linear time - O(n)
Algoritma ini membutuhkan waktu yang tumbuh sebanding dengan ukuran inputnya. Contohnya adalah perulangan sekuensial melalui array untuk mencari elemen tertentu.

* Linear time - O(n+m)
Algoritma ini menggambarkan kompleksitas waktu linear terhadap dua ukuran input yang berbeda, di mana "m" dan "n" adalah ukuran dari dua input yang berbeda. Dalam notasi ini, kompleksitas waktu tumbuh secara linier dengan jumlah dari kedua ukuran input tersebut. Contohnya saat kitaa memiliki dua himpunan data, seperti dua larik atau dua daftar, dengan ukuran yang berbeda, katakanlah m dan n, dan kita perlu melakukan operasi yang melibatkan kedua himpunan data tersebut. 

* Logarithmic time - O(log n)
Algoritma ini membutuhkan waktu yang meningkat secara logaritmik dengan ukuran inputnya. Contohnya adalah pencarian biner di dalam sebuah array terurut.

* Quadratic - O(n^2)
Algoritma ini membutuhkan waktu yang tumbuh sebanding dengan kuadrat dari ukuran inputnya. Contohnya adalah nested loop yang mengakses elemen dalam matriks 2D.

* Exponential Time - O(2^n)
Algoritma ini membutuhkan waktu yang meningkat secara eksponensial dengan ukuran inputnya. Contohnya adalah algoritma yang mencoba semua kemungkinan kombinasi, seperti algoritma brute-force untuk masalah NP-hard.

* Factorial Time - O(n!)
Algoritma ini mengindikasikan bahwa waktu eksekusi algoritma tumbuh sebanding dengan faktorial dari ukuran input (n). Algoritma dengan kompleksitas waktu O(n!) seringkali tidak praktis untuk digunakan, terutama untuk input yang besar, karena faktorial tumbuh sangat cepat. Contohnya adalah masalah permutasi lengkap yang mencoba semua kemungkinan permutasi dari n elemen.

## Time Limit
* n <= 1,000,000 dengan waktu yang diharapkan O(n^2)
* n <= 10,000 dengan waktu yang diharapkan O(n^2)
* n <= 500 dengan waktu yang diharapkan O(n^3)

## Space Complexity
Space complexity adalah konsep dalam analisis algoritma yang mengukur seberapa banyak ruang memori yang dibutuhkan oleh suatu algoritma untuk menyelesaikan masalah, sehubungan dengan ukuran inputnya. Space complexity mengukur seberapa efisien algoritma dalam penggunaan memori saat beroperasi. Space complexity juga memakai notasi O(Big O).


