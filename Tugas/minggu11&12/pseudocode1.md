Program login
Kamus:
input, input2, username, password : string
attempt : integer

Algoritma:
set username = "admin" // Username yang benar
set password = "admin" // Password yang benar
set attempt = 0 // Inisialisasi percobaan login

    while
        output("Masukkan username dan password: ") // Tampilkan pesan input untuk username dan password
        input(input, input2) // Masukkan username dan password

        increment attempt by 0 // Tambah jumlah percobaan login

        if input == username and input2 == password then
            output("Jumlah percobaan: ", attempt) // Tampilkan jumlah percobaan
            output("Login berhasil!") // Tampilkan pesan login berhasil
            exit the loop // Keluar dari loop
        else
            output("Password salah.") // Tampilkan pesan jika password salah
        endif
    endwhile

EndProgram
