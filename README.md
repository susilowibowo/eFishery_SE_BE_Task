# eFishery_SE_BE_Task
eFishery Software Engineer Back-end Task 

Petunjuk penggunaan :

- buat database mysql dengan nama "godb"

-(opsional) import file godb.sql ke database godb

- clone project, lalu jalankan applikasinya :
   - untuk Auth "go run Auth/main.go"
   - untuk Fetching "go run Fetching/main.go"

- masukan 4 angka untuk port-nya contoh: "8000","3000","8080"

- app ready to use

=======================================================

Menambah User :

- untuk pertama kali penggunaan, jalankan Auth, buka endpoint "/user" dengan method POST di Postman

- masukan form di body dengan key name, phone, role

- isi value tiap key lalu send


=====================================================

Auth Endpoint: 

- "/user"       [POST] : registrasi user 
- "/login"      [POST] : login untuk mendapatkan token, menggunakan phone dan password
- "/tokeninfo"  [POST] : menampilkan isi payload pada input token


Fetching Endpoint 

- "/login"       [POST] : login untuk mendapatkan token, menggunakan phone dan password
- "/showuser"    [GET]  : menampilkan info user yang ada di db
- "/showstorage" [GET]  : menampilkan seluruh hasil fetching + kurs harga dalam USD
<<<<<<< HEAD
- "/adminpage"   [POST]  : menampilkan data hasil fetching berdasarkan area_provinsi & weekly
=======
- "/adminpage"   [GET]  : menampilkan data hasil fetching berdasarkan area_provinsi & weekly
>>>>>>> a3fc87768b5662812f6cd92423d2c6d623079605
