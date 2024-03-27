# SERVICE UNTUK MEMBUKA APLIKASI SIDIK JARI BPJS

Rest api untuk membuka aplikasi sidik jari bpjs (windows desktop). Bisa digunakan untuk menghubungkan dengan aplikasi APM (anjungan pendaftaran mandiri) yang berbasis web.

spesifikasi

- Golang v 1.21.4
- Fiber
- [Keyboard event (github.com/micmonay/keybd_event)](https://github.com/micmonay/keybd_event) 


## Install
untuk membuat .exe

    go build -o FingerprintRestapi.exe main.go

letakan sejajar dengan aplikasi fingerprint BPJS (After.exe)

## Api spec
cek api

    curl --location --request POST 'http://localhost:3005/ping'

membuka aplikasi 

    curl --location 'http://localhost:3005/open' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "app_name": "After.exe",
    "username": "arif@gmail.com",
    "password": "Arif Kurniawan",
    "no_bpjs": "12112"
    }'

menutup aplikasi

    curl --location 'http://localhost:3005/close' \
    --header 'Content-Type: application/json' \
    --data '{
    "app_name": "After.exe"
    }'

