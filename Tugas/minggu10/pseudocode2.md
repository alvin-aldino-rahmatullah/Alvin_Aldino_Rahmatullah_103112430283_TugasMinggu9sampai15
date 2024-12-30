Program temperatur
Kamus:
    a, b, c, d, e : float64 

Algoritma:
    output("Masukkan lima nilai temperatur: ") // Tampilkan pesan input untuk lima temperatur
    input(a, b, c, d, e) // Masukkan lima nilai temperatur

    // Periksa stabilitas temperatur
    if a > b and b > c and c > d and d > e then
        output("stabil naik") // Temperatur stabil naik
    else if a < b and b < c and c < d and d < e then
        output("stabil turun") // Temperatur stabil turun
    else
        output("tidak stabil") // Temperatur tidak stabil
    endif

EndProgram
