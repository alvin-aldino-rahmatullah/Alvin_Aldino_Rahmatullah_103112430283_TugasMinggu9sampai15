Program profit
Kamus:
    bulan1, bulan2, total : float64 

Algoritma:
    output("Masukkan pendapatan bulan pertama: ") // Tampilkan pesan untuk input bulan pertama
    input(bulan1) // Masukkan pendapatan bulan pertama

    output("Masukkan pendapatan bulan kedua: ") // Tampilkan pesan untuk input bulan kedua
    input(bulan2) // Masukkan pendapatan bulan kedua

    // Perbandingan pendapatan
    if bulan2 > bulan1 then
        total <- bulan2 - bulan1 // Hitung peningkatan
        output("Pendapatan meningkat sejumlah ", total)
    else if bulan1 > bulan2 then
        total <- bulan1 - bulan2 // Hitung penurunan
        output("Penurunan sebesar ", total)
    else
        output("tetap") // Tidak ada perubahan pendapatan
    endif

EndProgram
