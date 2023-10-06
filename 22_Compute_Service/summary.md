# 22 Compute Service

## Deployment
Deployment adalah istilah dalam dunia teknologi dan pengembangan perangkat lunak yang merujuk pada proses mengambil aplikasi atau sistem yang telah dikembangkan dan menggerakkannya dari lingkungan pengembangan atau uji ke lingkungan produksi, di mana aplikasi tersebut dapat diakses oleh pengguna akhir.

Tujuan utama dari deployment adalah untuk membuat aplikasi atau sistem yang telah dikembangkan menjadi tersedia untuk pengguna akhir atau pelanggan. Ini melibatkan penggunaan berbagai metode dan alat untuk menginstal, mengkonfigurasi, menguji, dan memantau aplikasi atau sistem agar berjalan dengan baik di lingkungan produksi. Penyebaran layanan perangkat lunak dapat dilaukan ke berbagai platform, termasuk server, desktop, cloud, dan perangkat mobile.

## Strategi Deployment
### Big Bang Deployment
Big Bang Deployment adalah strategi di mana seluruh perubahan atau pembaruan perangkat lunak diterapkan secara bersamaan ke lingkungan produksi. Tujuannya adalah untuk menggulirkan perubahan secara cepat ke seluruh lingkungan dan menghindari kebingungan antara versi yang berbeda.

#### Keuntungan Big Bang Deployment:
* Kecepatan implementasi yang tinggi.
* Sederhana untuk dikelola.
* Konsistensi dalam versi aplikasi.

#### Kekurangan Big Bang Deployment:
* Risiko tinggi jika ada masalah atau bug.
* Tidak ada kemampuan rollback yang mudah.
* Dapat menyebabkan downtime yang signifikan.
* Pengujian terbatas sebelum peluncuran.
* Kesulitan identifikasi sumber masalah jika terjadi masalah.

### Rollout Deployment
Rollout Deployment adalah strategi di mana perubahan diterapkan secara bertahap ke lingkungan produksi. Ini dapat melibatkan beberapa tahap atau zona. Tujuannya adalah untuk mengendalikan risiko dengan memperkenalkan perubahan secara bertahap dan menguji setiap tahap.

#### Keuntungan Rollout Deployment:
* Mengurangi risiko dengan pengguliran bertahap.
* Memungkinkan deteksi masalah dengan cepat.
* Meminimalkan dampak negatif pada pengguna.
* Memungkinkan untuk menghentikan penyebaran jika terjadi masalah.

#### Kekurangan Rollout Deployment:
* Memerlukan lebih banyak waktu untuk menggulirkan perubahan secara penuh.
* Menambah kompleksitas manajemen dengan tahap-tahap penyebaran.
* Dapat memerlukan infrastruktur tambahan untuk tahap pengujian.

### Blue/Green Deployment
Blue/Green Deployment melibatkan dua lingkungan paralel (Blue dan Green). Versi perangkat lunak yang sedang berjalan (misalnya, Blue) tetap aktif, sementara versi baru (misalnya, Green) diterapkan di lingkungan yang terpisah. Tujuannya adalah untuk menghindari downtime dan memungkinkan beralih secara instan antara versi lama dan baru jika terjadi masalah.

#### Keuntungan Blue/Green Deployment:
* Tidak ada downtime karena versi lama dan baru berjalan bersamaan.
* Perubahan sanagat cepat
* Kemampuan pemutusan instan jika terjadi masalah.
* Memungkinkan rollback mudah ke versi sebelumnya.
* Cocok untuk aplikasi kritis yang tidak dapat mengalami downtime.

#### Kekurangan Blue/Green Deployment:
* Memerlukan infrastruktur ganda, yang dapat mahal.
* Memerlukan manajemen dan sinkronisasi yang cermat antara lingkungan Blue dan Green.
* Proses pengguliran dapat lebih rumit daripada beberapa strategi lainnya.

### Canary Deployment
Canary Deployment melibatkan pengguliran perubahan ke sebagian kecil dari lingkungan produksi (kelompok kanari) sebelum diimplementasikan secara luas. Tujuannya adalah untuk mendeteksi masalah atau bug pada tahap awal dan mengukur dampak perubahan secara bertahap.

#### Keuntungan Canary Deployment:
* Deteksi masalah dengan cepat pada tahap awal.
* Risiko rendah karena hanya sebagian kecil pengguna yang terpengaruh.
* Data pengguna nyata dapat digunakan untuk pengambilan keputusan.
* Kemampuan menguji perubahan dalam kondisi produksi.

#### Kekurangan Canary Deployment:
* Memerlukan manajemen yang cermat terhadap kelompok kanari.
* Diperlukan infrastruktur untuk mendukung penyebaran bertahap.
* Memerlukan pengawasan yang aktif terhadap kinerja kelompok kanari.
* Tidak cocok untuk semua jenis aplikasi atau perubahan.

## Process Deployment
* Development Environment: Tempat pengembang membuat kode perangkat lunak.
* Test Environment: Lingkungan pengujian perangkat lunak.
* UAT Environment: Lingkungan pengujian oleh pengguna akhir sebelum produksi.
* Production Environment: Lingkungan produksi di mana perangkat lunak berjalan untuk pengguna akhir.

## Git Procces Secara Sederhana
* Clone Repository (Kloning Repositori): Membuat salinan remote repository ke komputer lokal Anda.
* Workspace (Ruang Kerja): Tempat di mana Anda bekerja pada proyek, membuat dan mengedit file.
* Local Repository (Repositori Lokal): Repositori Git di komputer lokal yang mencerminkan remote repository.
* Commit (Mengcommit Perubahan): Mengabadikan perubahan dari workspace ke repositori lokal dengan pesan deskripsi.
* Push to Remote (Push ke Remote): Mengirim perubahan dari repositori lokal ke remote repository.
* Pull from Remote (Pull dari Remote): Mengambil perubahan dari remote repository ke repositori lokal Anda.

## CLI & SHELL Scripts
CLI adalah antarmuka teks yang digunakan untuk berinteraksi dengan sistem, sementara shell script adalah kumpulan perintah yang dijalankan dalam lingkungan shell untuk mengotomatisasi tugas atau proses tertentu.

## Deployment Compute Engine / EC2
```
* Cara Membuat Instance VM
https://cloud.google.com/compute/docs/instances/create-start-instance
https://docs.aws.amazon.com/id_id/iot/latest/developerguide/creating-a-virtual-thing.html

* Cara Membuat Security Group or Firewall
https://cloud.google.com/firewall/docs/using-firewalls
https://docs.aws.amazon.com/id_id/AWSEC2/latest/UserGuide/working-with-security-groups.html

* Cara Create SSH Key
https://cloud.google.com/compute/docs/connect/create-ssh-keys#windows-10-or-
https://docs.aws.amazon.com/id_id/ground-station/latest/ug/create-ec2-ssh-key-pair.html

* Cara Add SSH keys to VMs
https://cloud.google.com/compute/docs/connect/add-ssh-
https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html

```

## SSH Remote ke VM
```
ssh -i </directory/namafilessh> <username-server>@<public-ipv4>
```

## Cara Copy Folder Windows to Ubuntu
```
cp -r "<location-folder>" ~/
```

## Cara Transfer File/Folder ke Server
```
* Transfer file
scp -i </direktori/ssh-key-private> </direktori/nama-file-transfer> <username>@<public-ip-server>:<goal-location-server>

* Transfer folder
scp -i </direktori/ssh-key-private> -r </direktori/nama-file-transfer> <username>@<public-ip-server>:<goal-location-server>
```

## Cek Files
```
ls
```

## Cara Membuat Database GCP
Ketika Anda membuat SQL database atau server, penting untuk mengatur konfigurasi jaringan dengan benar agar dapat diakses dari server lain
```
https://cloud.google.com/sql/docs/mysql/create-manage-databases
https://cloud.google.com/sql/#section-1
https://docs.aws.amazon.com/id_id/AWSEC2/latest/UserGuide/option2-task2-create-rds-database.html
```

