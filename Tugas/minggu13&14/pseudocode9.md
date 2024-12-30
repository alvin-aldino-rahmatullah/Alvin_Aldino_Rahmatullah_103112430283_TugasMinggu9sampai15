Program Polab

Kamus:
k, j, n : integer

Algoritma:
output("Masukkan bilangan positif: ") // Tampilkan pesan untuk input n
input(n) // Masukkan nilai n

    for j <- 1 to n do // Perulangan untuk setiap baris
        for k <- 1 to n do // Perulangan untuk setiap kolom
            if j == 1 or j == n or k == 1 or k == n then
                output(j, " ") // Cetak angka baris jika berada di tepi
            else
                output("  ") Cetak spasi untuk bagian dalam
            endif
        endfor
        output() // Pindah ke baris berikutnya
    endfor

EndProgram
