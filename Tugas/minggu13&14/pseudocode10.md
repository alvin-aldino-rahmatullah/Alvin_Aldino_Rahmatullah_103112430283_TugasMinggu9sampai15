Program Polax

Kamus:
n, i, j : integer

Algoritma:
output("Masukkan bilangan positif: ") // Tampilkan pesan untuk input n
input(n) // Masukkan nilai n

    if n mod 2 == 0 then
        output("Input harus bilangan positif")// Validasi input harus bilangan ganjil
        return
    endif

    for i <- 0 to n-1 do // Perulangan untuk setiap baris
        for j <- 0 to n-1 do // Perulangan untuk setiap kolom
            if i == j or i + j == n - 1 then
                output(j+1, " ") // Cetak angka jika posisi pada diagonal utama atau sekunder
            else
                output("  ") // Cetak spasi untuk bagian lainnya
            endif
        endfor
        output() // Pindah ke baris berikutnya
    endfor

EndProgram
