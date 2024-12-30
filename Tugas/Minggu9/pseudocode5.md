Program bola
Kamus:
    a, b, c, d : integer  // Jumlah gol empat tim
    max, min : integer    // Gol terbanyak dan tersedikit
Algoritma:
    output("Masukkan jumlah gol empat tim: ")  // Tampilkan pesan untuk input
    input(a, b, c, d)                          // Masukkan jumlah gol keempat tim

    // Menentukan gol terbanyak
    if a > b then
        max <- a
    else
        max <- b
    endif

    if c > max then
        max <- c
    endif

    if d > max then
        max <- d
    endif

    // Menentukan gol tersedikit
    if a < b then
        min <- a
    else
        min <- b
    endif

    if c < min then
        min <- c
    endif

    if d < min then
        min <- d
    endif

    output("Gol terbanyak:", max)      // Tampilkan gol terbanyak
    output("Gol tersedikit:", min)    // Tampilkan gol tersedikit
EndProgram
