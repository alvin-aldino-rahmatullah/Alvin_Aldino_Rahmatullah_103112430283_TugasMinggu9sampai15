Program terbesar

Kamus:
digit, maxDigit : integer

Algoritma:
output("Masukkan bilangan positif: ") // Tampilkan pesan untuk input digit
input(digit) // Masukkan nilai digit

    maxDigit <- -1 // Inisialisasi digit terbesar

    while digit > 0 do // Perulangan untuk memproses setiap digit
        if digit mod 10 > maxDigit then
            maxDigit <- digit mod 10 // Perbarui nilai digit terbesar
        endif
        digit <- digit div 10 // Buang digit terakhir
    endwhile

    output(maxDigit) // Tampilkan digit terbesar

EndProgram
