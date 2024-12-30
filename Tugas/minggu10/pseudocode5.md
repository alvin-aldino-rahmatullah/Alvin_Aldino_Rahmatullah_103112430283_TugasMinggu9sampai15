Program akhirtahun
Kamus:
totalBelanja, diskon, cashback : integer
buatKartu : boolean

Algoritma:
output("Masukkan total belanja : ")
input(totalBelanja)
output("Apakah mau membuat kartu? (true/false): ")
input((buatKartu)

    // Pengecekan status kartu
    if buatKartu then
        output("Kartu? true")
    else
        output("Kartu? false")
    endif

    // Pengecekan dan penerapan diskon
    if totalBelanja >= 100000 then
        output("Diskon? true")
        diskon := totalBelanja * 10 / 100
        totalBelanja := totalBelanja - diskon
        output("Diskon yang diterima: Rp.", diskon)
    else
        output("Diskon? false")
    endif

    // Pengecekan dan penerapan cashback
    if totalBelanja >= 200000 then
        output("Cashback? true")
        cashback := 75000
        totalBelanja := totalBelanja - cashback
        output("Cashback yang diterima: Rp.", cashback)
    else
        output("Cashback? false")
    endif

    // Menampilkan total belanja setelah diskon dan cashback
    output("Total Belanja setelah diskon dan cashback: Rp.", totalBelanja)

EndProgram
