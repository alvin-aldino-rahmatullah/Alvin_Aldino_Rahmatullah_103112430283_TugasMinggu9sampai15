Program segitiga
Kamus:
    a, b, c : integer 

Algoritma:
    output("Masukkan Jumlah sisi segitiga : ") // Tampilkan pesan input untuk sisi 
    input(a,b,c) // Baca input dari pengguna 

    // Cek jenis segitiga
    if a = b and b = c then
        output("Segitiga sama sisi") // Semua sisi sama panjang
    else if a ! b and b ! c and a ! c then
        output("Segitiga sembarang") // Semua sisi berbeda panjang
    else
        output("Segitiga sama kaki") // Dua sisi sama panjang
    endif

EndProgram
