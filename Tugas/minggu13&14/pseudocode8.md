Program Polaa

Kamus:
k, j, n : integer

Algoritma:
output("Masukkan bilangan positif: ") // Tampilkan pesan untuk input n
input(n) // Masukkan nilai n

    for k <- 1 to n do // Perulangan untuk setiap baris
        for j <- 1 to n do // Perulangan untuk setiap kolom
            output(j, " ") // Cetak nilai j di baris tersebut
        endfor
        output("") // ke baris berikutnya
    endfor

EndProgram
