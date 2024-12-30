Program kopi
Kamus:
n, m, x, y, cangkir : integer
Algoritma:
output("Masukkan jumlah gula, kopi, kebutuhan gula per cangkir, kebutuhan kopi per cangkir (pisahkan dengan spasi): ")
input(n, m, x, y) // Masukkan nilai n, m, x, y

    cangkir <- 0  // Inisialisasi jumlah cangkir yang dapat dibuat

    while n >= x and m >= y do
        n <- n - x  // Kurangi gula untuk membuat satu cangkir
        m <- m - y  // Kurangi kopi untuk membuat satu cangkir
        cangkir <- cangkir + 1  // Tambah jumlah cangkir
    endwhile

    output("Jumlah cangkir kopi yang bisa dibuat: ", cangkir)

EndProgram
