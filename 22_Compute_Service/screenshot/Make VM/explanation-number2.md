# Penjelasan Soal Nomor 2
Di Google Cloud Platform (GCP), SSH remote ke sebuah virtual machine (VM) biasanya dilakukan dengan menggunakan key pair SSH.

## Make Key Pair SSH in Local PC
```
ssh-keygen -t ecdsa
```

## Penjelasan Key
* Private Key: Private key adalah kunci rahasia yang harus disimpan dengan sangat aman di komputer lokal. Ini adalah kunci yang akan digunakan untuk mengenkripsi dan mendekripsi komunikasi SSH antara komputer dan VM di GCP/AWS. Private key tidak boleh dibagikan atau diunggah ke tempat yang tidak aman.

* Public Key: Public key adalah kunci yang bersifat publik dan dapat ditemukan di dalam metadata VM di GCP/AWS. Public key ini digunakan untuk mengidentifikasi sebagai pengguna yang sah saat mencoba SSH ke VM. Public key tidak memiliki kemampuan untuk mengenkripsi atau mendekripsi pesan, itu hanya digunakan untuk verifikasi.

## Add Key Public in VM
Biasanya berada pada folder dibawah ini, gunakan yang .pub (key public). Baca isi key nya dan paste ke VM

```
~/.ssh/id_rsa.pub
```

## SSH Remote Menggunakan Key Pair
```
ssh -i /path/to/private-key username@public-ip
```

## Password in GCP/AWS
Secara umum, di GCP/AWS, password tidak digunakan untuk otentikasi SSH ke VM, kecuali kita secara khusus mengaktifkan otentikasi berbasis password di dalam VM. Namun, disarankan untuk menggunakan key pair SSH untuk keamanan yang lebih baik.

## Cara Atur Password
Namun di AWS (Amazon Web Services) dan GCP (Google Cloud Platform) kita dapat mengaktifkan otentikasi berbasis password untuk SSH dengan langkah-langkah berikut:
* Atur password untuk pengguna SSH yang ada di dalam VM Anda.
* Ubah konfigurasi SSH di VM untuk mengaktifkan otentikasi berbasis password.
* Restart layanan SSH di VM untuk menerapkan perubahan.




