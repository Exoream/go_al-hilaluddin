# 19 Unit Testing

## Unit Testing
Unit testing adalah jenis pengujian perangkat lunak yang fokus pada pengujian unit individu atau komponen kecil dalam kode. Unit dalam konteks ini adalah bagian terkecil dari program yang dapat diuji secara terisolasi. Biasanya, unit ini adalah fungsi atau metode individu, atau bahkan bagian lebih kecil dari fungsi tersebut.

## Tujuan Penggunaan Unit Testing
```
~ Mengidentifikasi Masalah Dengan Cepat
~ Meminimalkan Dampak Perubahan
~ Dokumentasi Hidup
~ Memudahkan Refaktor
~ Pengujian Otomatis
~ Membantu dalam Debugging
~ Peningkatan Kualitas Kode
```

## Level Testing
```
Unit Testing adalah pengujian pada level terendah, yang menguji komponen kode secara terisolasi. 

Integration Testing menguji interaksi antara unit-unit yang telah diuji sebelumnya, memastikan bahwa komponen-komponen ini berfungsi dengan baik ketika digabungkan. 

UI Testing berkaitan dengan pengujian antarmuka pengguna aplikasi, memeriksa apakah pengguna dapat berinteraksi dengan aplikasi secara efektif melalui elemen-elemen UI.
```

## Framework Unit Test
```
Go
~ Testify
```

```
JavaScript
~ Jest
~ Mocha
~ Jasmine
```

## Structure
Structure unit testing biasanya ada 2 yaitu memusatkan file pengujian di dalam test folder dan menyimpan file pengujian bersama dengan file produksi

## Hal yang Tidak Perlu di Test
```
~ 3rd party modules
~ Database
~ 3rd party API
~ Object class yang ditest di tempat lain
```

## Mock Object
Mock Object adalah salah satu konsep dalam pengujian perangkat lunak yang digunakan untuk menggantikan komponen-komponen nyata (seperti objek atau layanan) dengan objek palsu yang dapat dikontrol dengan lebih baik selama pengujian. 

## Coverage
Coverage dalam unit testing merujuk pada sejauh mana kode sumber diuji oleh tes unit. Ini memberikan ukuran seberapa banyak kode yang dieksekusi selama pengujian dan membantu mengidentifikasi sejauh mana kode yang belum diuji. Coverage report terdiri dari function, statement, branch, dadn lines. Sedangkan report memiliki format CLI, XML, HTML, dan LCOV.

## New Test In Golang
```
~ Nama file harus "namafile_test.go"
~ Location file boleh sama folder, sama package atau sama folder, berbeda package
```

## Run Testing
```
go test ./nama_folder - cover
```

```
go test ./nama_folder -coverprofile=cover.out && go tool cover -html=cover.out
```



