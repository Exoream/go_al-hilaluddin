# CICD

## Continous Integration
Continuous Integration (CI) adalah proses otomatis yang digunakan untuk mengintegrasikan berbagai perubahan kode dari berbagai sumber potensial ke dalam repositori kode bersama secara teratur dan sering. Proses ini dapat mencakup kompilasi, pengujian, dan pemeriksaan otomatis lainnya untuk memastikan bahwa perubahan kode tersebut bekerja dengan baik dan tidak mengganggu fungsionalitas yang ada.

## Siklus Continous Integration
* Check-in Change: Tim pengembang mengirimkan perubahan kode ke repositori bersama.
* Fetch Change: Sistem CI mengambil perubahan terbaru dari repositori.
* Build: Kode yang diubah dikompilasi menjadi aplikasi yang dapat dijalankan.
* Test: Aplikasi yang dikompilasi diuji dengan berbagai jenis pengujian, termasuk pengujian unit dan pengujian fungsional.
* Fail or Success: Hasil pengujian ditentukan, dan apakah perubahan kode berhasil atau gagal diidentifikasi.
* Notify Success or Failure: Tim atau pemangku kepentingan diberi tahu tentang hasil pengujian, apakah perubahan berhasil atau gagal.

## Continuous Deployment dan Continuous Delivery
Continuous Deployment dan Continuous Delivery adalah langkah yang lebih maju dari Continuous Integration (CI) dalam praktik pengembangan perangkat lunak. Dalam kedua praktik ini, tujuannya adalah untuk secara otomatis menyebarkan perubahan kode ke lingkungan produksi atau pengujian, tetapi perbedaan utamanya adalah pada sejauh mana otomatisasi ini dilakukan dan apakah ada campur tangan manusia sebelum perubahan tersebut diterapkan ke lingkungan produksi.

Continuous Deployment: Setiap perubahan kode yang berhasil diuji secara otomatis langsung diterapkan ke lingkungan produksi tanpa campur tangan manusia. Pengiriman otomatis ke produksi.

Continuous Delivery: Setiap perubahan kode yang berhasil diuji secara otomatis siap untuk diimplementasikan ke produksi, tetapi penerapannya memerlukan campur tangan manusia. Pengiriman ke produksi dilakukan secara manual.

## Siklus Continuous Deployment dan Continuous Delivery
### Continuous Deployment
* Unit Test: Pengujian unit untuk memeriksa bagian-bagian kecil kode. (Auto)
* Platform Test: Pengujian aplikasi pada platform yang sesuai. (Auto)
* Deliver to Staging: Pengiriman ke lingkungan staging untuk pengujian lebih lanjut. (Auto)
* Application Acceptance Test: Pengujian oleh pengguna atau pemangku kepentingan. (Auto)
* Deploy to Production: Implementasi otomatis ke produksi. (Auto)
* Post Deploy Tests: Pengujian pasca-implementasi. (Auto)

### Continuous Delivery
* Unit Test: Pengujian unit untuk memeriksa bagian-bagian kecil kode. (Auto)
* Platform Test: Pengujian aplikasi pada platform yang sesuai. (Auto)
* Deliver to Staging: Pengiriman ke lingkungan staging untuk pengujian lebih lanjut. (Auto)
* Application Acceptance Test: Pengujian oleh pengguna atau pemangku kepentingan. (Auto)
* Deploy to Production: Implementasi ke produksi. (Manual)
* Post Deploy Tests: Pengujian pasca-implementasi. (Auto)

## Contoh Tools Yang Digunakan
#### Code and Commit
* Git
* Github
* Bitbucket

#### Build and Config
* Docker
* Maven
* Ansible
* Jenkins

#### Scan and Test
* JUnit
* SonarQube
* Jenkins Pipeline
* Selenium

#### Release
* Serena
* XL Release
* uDeploy

#### Deploy
* Kubernetes
* Docker
* Google Cloud Engine
* VMware

## Kesimpulan
Continuous Integration (CI) dan Continuous Deployment/Delivery (CD) adalah praktik-praktik penting dalam pengembangan perangkat lunak modern. CI memungkinkan pengujian dan integrasi kode yang lebih cepat dan otomatis, sementara CD memungkinkan pengiriman perangkat lunak yang cepat dan konsisten ke lingkungan produksi. Alat-alat seperti Jenkins, Docker, Maven, Ansible, dan banyak lainnya digunakan dalam tahap-tahap berbeda dari alur kerja CI/CD, termasuk code and commit, build and config, scan and test, release, dan deploy. Keseluruhan, praktik-praktik ini membantu organisasi untuk menghasilkan perangkat lunak berkualitas tinggi, aman, dan responsif dengan efisiensi yang lebih besar dalam menghadapi perubahan pasar yang cepat.

## Cara CI/CD Di Github Action
* Buat direktori .github/workflows di repositori GitHub Anda jika belum ada.
* Buat berkas konfigurasi YAML untuk CI dan CD di direktori tersebut.
* Konfigurasi langkah-langkah CI, seperti pengujian dan pemeriksaan kode, dalam berkas YAML.
* Konfigurasi langkah-langkah CD, seperti proses pengiriman, dalam berkas YAML.
* Simpan perubahan dan komit ke repositori Anda untuk mengaktifkan GitHub Actions.
Note : Jika menggunakan secret variable, tambahkan kedalam github repository nya dibagian settings dan menu actions secrets and variables.

## Kubernetes
Kubernetes adalah platform orkestrasi wadah yang open-source yang digunakan untuk mengelola, merancang, dan mengimplementasikan aplikasi berbasis kontainer. Ini membantu dalam otomatisasi implementasi, penskalaan, dan manajemen aplikasi yang berjalan dalam kontainer di berbagai lingkungan, termasuk pusat data lokal dan lingkungan cloud.

## Alasan Membutuhkan Container Orchestration System
Container orchestration system, seperti Kubernetes, dibutuhkan untuk mengelola dan mengotomatisasi aplikasi yang berjalan dalam kontainer di berbagai lingkungan. Mereka mempermudah penskalaan, pengaturan, pemantauan, dan pemulihan aplikasi, serta memastikan ketersediaan dan keamanan. Orkestrasi juga mengotomatisasi implementasi, distribusi beban, dan pembaruan aplikasi, membantu tim pengembangan mengelola infrastruktur yang kompleks secara efisien.

## Komponen Utama Kubernetes Cluster Diagram
* Master Node: Otak cluster yang mengendalikan operasi. Ini termasuk API Server, Etcd (penyimpanan data), Kube Scheduler (penjadwal tugas), dan Kube Controller Manager (pengelola kontroler).
* Node (Worker Node): Mesin di mana kontainer aplikasi berjalan. Setiap node memiliki Kubelet (komunikasi dengan master), * Kube Proxy (pengaturan jaringan), dan runtime kontainer.
* Pods: Unit terkecil dalam Kubernetes tempat kontainer berjalan.
* Service: Abstraksi lapisan jaringan untuk menghubungkan pod-pod.
* Ingress: Mengelola lalu lintas eksternal ke dalam cluster.
* Namespace: Digunakan untuk mengorganisasi dan mengisolasi objek-objek Kubernetes dalam cluster.

## Konsep Penting Kubernetes
* Service Discovery dan Load Balancing: Kubernetes membantu aplikasi menemukan layanan dan mendistribusikan lalu lintas secara otomatis.
* Horizontal Scaling: Aplikasi dapat dengan mudah ditingkatkan atau dikurangi berdasarkan permintaan.
* Automated Rollouts dan Rollback: Pembaruan aplikasi bisa dilakukan tanpa waktu henti dan otomatis di-rollback jika ada masalah.
* Secret and Configuration Management: Kubernetes aman menyimpan dan mengelola konfigurasi dan data rahasia aplikasi.

## Hal Yang Tidak Dapat Dilakukan Kubernetes
* Tidak Membatasi Jenis Aplikasi yang Didukung
* Tidak Mengimplementasikan Kode Sumber dan Tidak Membangun Aplikasi
* Tidak Menyediakan Layanan Tingkat Aplikasi, Seperti Middleware
* Tidak Mengatur Solusi Logging, Monitoring, atau Alerting

## Kubernetes Workloads
* Namespace: Cara untuk mengorganisasi dan mengisolasi objek-objek Kubernetes dalam cluster.
* Pod: Unit terkecil dalam Kubernetes yang berisi satu atau beberapa kontainer.
* Labels: Metadata yang digunakan untuk mengkategorikan dan mengidentifikasi objek Kubernetes.
* Deployment: Deployment digunakan untuk mengelola aplikasi dengan replika pod, mendukung rolling updates, dan penskalaan otomatis.
* Rolling Update: Metode pembaruan aplikasi tanpa waktu henti.
* Secret: Objek yang digunakan untuk menyimpan data rahasia seperti kata sandi.
* Service: Abstraksi lapisan jaringan untuk menghubungkan pod dalam cluster.
* Ingress: Kontroler yang mengatur lalu lintas eksternal ke layanan dalam cluster.

## Domain Name System
Domain Name System (DNS) adalah sistem yang digunakan untuk menerjemahkan alamat IP numerik ke nama domain yang lebih mudah diingat.

