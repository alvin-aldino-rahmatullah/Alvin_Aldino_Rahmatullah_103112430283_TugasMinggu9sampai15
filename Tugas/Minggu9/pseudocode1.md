Program Tamasya
Kamus:
    n, mobil : integer
Algoritma:
    output("berapa orang yang akan naik mobil?")  // Tampilkan pesan
    input(n)  // Masukkan jumlah orang

    if n mod 7 == 0 then
        mobil <- n div 7  // Jika jumlah orang kelipatan 7, hitung mobil
    else
        mobil <- (n div 7) + 1  // Jika tidak, tambah 1 mobil
    endif

    output("jumlah mobil anda adalah", mobil)  // Tampilkan jumlah mobil
EndProgram
