# SERVICE UNTUK MEMBUKA APLIKASI SIDIK JARI BPJS

Rest api untuk membuka aplikasi sidik jari bpjs (windows desktop). Bisa digunakan untuk menghubungkan dengan aplikasi APM (anjungan pendaftaran mandiri) yang berbasis web.

spesifikasi

- Golang v 1.21.4
- [Fiber](https://adminlte.io/)
- [Windows api wrapper (github.com/lxn/win)](https://github.com/lxn/win) 
- [Keyboard event (github.com/micmonay/keybd_event)](https://github.com/micmonay/keybd_event) 

## Api
membuka aplikasi (foreground)

    <?php

        $curl = curl_init();

        curl_setopt_array($curl, [
        CURLOPT_PORT => "3001",
        CURLOPT_URL => "http://localhost:3001/open",
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_ENCODING => "",
        CURLOPT_MAXREDIRS => 10,
        CURLOPT_TIMEOUT => 30,
        CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
        CURLOPT_CUSTOMREQUEST => "POST",
        CURLOPT_POSTFIELDS => "{\"app_name\":\"Aplikasi Registrasi Sidik Jari\",\"no_indetitas\":\"ini string\"}",
        ]);

        $response = curl_exec($curl);
        $err = curl_error($curl);

        curl_close($curl);

        if ($err) {
        echo "cURL Error #:" . $err;
        } else {
        echo $response;
        }

menutup aplikasi (minimize)

    <?php

    $curl = curl_init();

    curl_setopt_array($curl, [
    CURLOPT_PORT => "3001",
    CURLOPT_URL => "http://localhost:3001/close",
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_ENCODING => "",
    CURLOPT_MAXREDIRS => 10,
    CURLOPT_TIMEOUT => 30,
    CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
    CURLOPT_CUSTOMREQUEST => "POST",
    CURLOPT_POSTFIELDS => "{\"app_name\":\"Aplikasi Registrasi Sidik Jari\"}",
    ]);

    $response = curl_exec($curl);
    $err = curl_error($curl);

    curl_close($curl);

    if ($err) {
    echo "cURL Error #:" . $err;
    } else {
    echo $response;
    }


## Build
untuk membuat .exe

    go build -o service-sidikjari-bpjs.exe main.go

jalankan di pc yang terpasang aplikasi sidik jari bpjs.