# 21 Docker

## Docker
Docker adalah platform perangkat lunak yang digunakan untuk mengelola dan menjalankan aplikasi dalam wadah (container). Wadah adalah lingkungan terisolasi yang memungkinkan Anda untuk menjalankan aplikasi dan semua dependensinya dalam satu unit yang dapat dijalankan di berbagai lingkungan komputasi, seperti server fisik, mesin virtual, atau cloud.

## Container
Container adalah unit perangkat lunak yang berisi aplikasi beserta semua dependensinya, termasuk perpustakaan, konfigurasi, dan kode. Container mengemas aplikasi dan lingkungan yang diperlukan untuk menjalankannya dalam unit terisolasi yang dapat dijalankan di berbagai lingkungan komputasi, seperti server fisik, mesin virtual, atau cloud.


## Perbedaan Utama antara Container dan Virtual Machine
#### Container
* Container adalah cara untuk mengisolasi aplikasi dan dependensinya dari sistem host. Abstraksi ini terjadi pada lapisan aplikasi.

* Container lebih efisien dalam penggunaan ruang dibandingkan dengan mesin virtual (VM) karena mereka menggunakan kernel sistem operasi host. Ini berarti mereka menggunakan sumber daya yang lebih sedikit.

* Container cocok untuk menjalankan lebih dari satu aplikasi pada satu host. Mereka bisa menangani beberapa aplikasi dengan efisien, sehingga mengurangi penggunaan sumber daya.

* Karena sifatnya yang ringan, container memerlukan lebih sedikit VM dan sistem operasi untuk menjalankan beberapa aplikasi, yang mengurangi overhead.
Mesin Virtual (VM):

#### Mesin Virtual (VM)
* VM mengabstraksi perangkat keras fisik, menciptakan lingkungan virtual yang mencakup salinan lengkap dari sistem operasi dan aplikasi.

* Setiap VM beroperasi sebagai unit independen dengan sistem operasi penuh yang terpisah. Ini memberikan isolasi yang kuat tetapi juga mengakibatkan penggunaan sumber daya yang lebih tinggi.

* VM bisa lebih lama untuk boot dibandingkan dengan container karena mereka harus memulai seluruh sistem operasi bersama dengan aplikasinya. Ini mengakibatkan waktu mulai yang lebih lama.


## Docker Basic
* Image: Paket yang berisi aplikasi dan dependensinya, digunakan sebagai dasar untuk membuat container.
* Container: Lingkungan terisolasi tempat aplikasi berjalan, dibuat dari image Docker.
* Docker Engine: Inti Docker yang menjalankan container dan menyediakan antarmuka command line.
* Registry: Tempat penyimpanan image Docker, seperti Docker Hub.
* Control Panel: Alat grafis untuk mengelola dan memantau kontainer dan image Docker.

## Docker Infrastruktur
* Client: Docker client adalah perangkat lunak yang memungkinkan berinteraksi dengan Docker Engine. Kita menggunakan Docker client untuk menjalankan perintah dan mengelola kontainer serta image. 

* Docker Host: Docker host adalah komputer fisik atau virtual tempat kontainer Docker sebenarnya berjalan. Docker Engine diinstal di Docker host dan bertanggung jawab atas menjalankan, mengelola, dan mengisolasi kontainer. Setiap host dapat menjalankan beberapa kontainer.

* Registry: Registry adalah tempat penyimpanan image Docker. Ini adalah layanan atau repositori tempat kita dapat mencari, mengunduh, mengunggah, dan berbagi image. Docker Hub adalah registry publik yang populer, tetapi kita juga dapat mengonfigurasi registry pribadi jika diperlukan.

## Tahapan Permulaan Menggunakan Docker
* Instal Docker di komputer (Windows, macOS, atau Linux).
* Pelajari konsep dasar Docker seperti image, container, Dockerfile, dan registry.
* Cari image Docker yang sesuai di Docker Hub atau registry lainnya.
* Buat Dockerfile jika tidak ada image yang sesuai. Ini berisi instruksi untuk membuat image.
* Gunakan docker build untuk membuat image dari Dockerfile.
* Jalankan container dari image dengan docker run.
* Kembangkan aplikasi dalam container.
* Uji aplikasi dalam lingkungan kontainer.
* Optimalkan sumber daya dan pantau kinerja.
* Dokumentasikan langkah-langkah dan hasil proyek.
* Deploy aplikasi ke lingkungan produksi dengan alat orkestrasi jika diperlukan.

## Syntax Docker
* FROM: Menentukan image dasar yang akan digunakan sebagai titik awal pembuatan image baru.
* RUN: Menjalankan perintah-perintah selama proses pembuatan image untuk melakukan instalasi atau konfigurasi.
* ENV: Mengatur variabel lingkungan dalam container.
* ADD: Menyalin file dari host ke dalam image.
* COPY: Menyalin file dari host ke dalam image (mirip dengan ADD, tetapi lebih sederhana).
* WORKDIR: Mengatur direktori kerja default dalam container.
* ENTRYPOINT: Menentukan perintah utama yang akan dijalankan saat kontainer dimulai.
* CMD: Menyediakan argumen default yang dapat diabaikan ketika kontainer dijalankan bersamaan dengan ENTRYPOINT.
* EXPOSE: Mendefinisikan port yang akan dibuka untuk koneksi dari luar saat kontainer dijalankan.
* USER: Mengatur pengguna yang akan digunakan saat menjalankan perintah dalam kontainer.
* VOLUME: Mendefinisikan volume untuk menyimpan data persisten dalam kontainer.
* ARG: Mendefinisikan argumen yang dapat digunakan saat membangun image Docker.
* LABEL: Menambahkan metadata ke image Docker.
* HEALTHCHECK: Mendefinisikan periksaan kesehatan untuk pemantauan aplikasi.
* SHELL: Mengganti shell default yang digunakan dalam perintah RUN.

## Docker Volume
Docker volume adalah mekanisme yang digunakan dalam Docker untuk mengelola penyimpanan data persisten yang terkait dengan kontainer. Mereka digunakan untuk menyimpan data yang bertahan bahkan setelah kontainer dihentikan atau dihapus.

## Docker Network
Docker Network adalah fitur dalam Docker yang memungkinkan Anda untuk mengatur dan mengelola jaringan antara kontainer Docker dan antara kontainer dan host fisik. Docker Network memungkinkan komunikasi antara kontainer yang berjalan di lingkungan yang diisolasi.

## Cara Instal Docker di Ubuntu
#### Set up Docker's Apt repository
```
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

# Add the repository to Apt sources:
echo \
  "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
```

#### Install the Docker packages.
```
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

## Cara Instal Compose Docker
```
sudo apt-get update
sudo apt-get install docker-compose-plugin
```

## Cek Docker Sudah Terinstal
```
docker --version
docker compose version
```

## Contoh Code Dockerfile
```
FROM golang:1.21.0-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o hilal .

EXPOSE 8000

CMD ["/app/hilal"]
```

## Docker Build
```
docker build -t <nama-image>:<tag> .
```

## Show Images
```
docker images
docker images list
```

## Delete Images
```
docker rmi <image-id>
docker rmi <image-name>
```

## Create Container
```
docker run -p 8080:8080 --name <container-name> <image-name>:<tag>

* with env
docker run -d -p 80:80 -e DBUSER=root -e DBPASS=123 -e DBHOST=127.0.0.1 -e DBPORT=3306 -e DBNAME=docker --name <container-name> <image-name>:<tag>
``` 

## Show Container
```
* yang berjalan
docker ps

* seluruh container termasuk berhenti
docker ps -a
```

## Star Container
```
docker start <container-name>
```

## Stop Container
```
docker stop <container-name>
```

## Remove Container
```
docker rm <container-name>
docker rm <container-id>
```

## Logs Container
```
docker logs <container-name>
```

## Push Image to Registry
```
docker login -u <username-dockerhub>
masukkan password
docker tag <image-name>:<tag> <username-dockerhub>/<image-name>
docker push <username-dockerhub>/<image-name>
```

## Pull Repository
```
docker pull <image-name>
```

## Create, Remove and Add Container Volume
```
docker create volume <name_volume>
docker volume rm <name_volume>
docker container create --name <name_container> \ --mount "type=volume, source=<name_volume>, dst=<folder_destination>"
```

## Create, Delete and List Network
```
docker network create <name_network>
docker network rm <name_network>
docker ls
```

## Container to Network
```
* regist container to network
docker container create --name <container_name> \ --network <name_network>

* delete container from network
docker network disconect <name_network> <container_name>
```
## Check IP Container
```
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' <container-name>
```

## Instal MySQL
```
* pakai password
docker run --name <container-name> -e MYSQL_ROOT_PASSWORD=password -d <image-name>:<tag>

* tanpa password
docker run --name <container-name> -e MYSQL_ALLOW_EMPTY_PASSWORD=true -d mysql:<tag>

* allow all akses
docker run --name <container-name> -e MYSQL_ROOT_PASSWORD=password -e MYSQL_ROOT_HOST=% -d mysql:<tag>
```

## Cara Masuk Database
```
docker exec -it <container-name> bash
```

