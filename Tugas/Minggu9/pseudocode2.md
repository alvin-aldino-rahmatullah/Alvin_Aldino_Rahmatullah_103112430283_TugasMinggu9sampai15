Program diskon
Kamus:
uang, total : float64
assemen : boolean
Algoritma:
output("Masukkan uang kamu: ") // Tampilkan pesan untuk input uang
input(uang) // Masukkan nilai uang

    output("Apakah kamu sedang asesmen (true/false)? ")  // Tampilkan pesan untuk input asesmen
    input(assemen)                   // Masukkan nilai boolean untuk asesmen

    if assemen == true then
        total <- uang * 0.65         // Hitung total setelah diskon 35%
        output("Kamu memperoleh potongan 35% dari ", uang, " menjadi ", total)
    else
        output("Kamu tidak sedang asesmen jadi tidak ada diskon")
    endif

EndProgram
