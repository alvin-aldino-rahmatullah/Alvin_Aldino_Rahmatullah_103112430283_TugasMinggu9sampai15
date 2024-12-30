program cari

kamus
x, n, digit: integer
found: boolean

algoritma
output("Masukkan bilangan x: ") // Tampilkan pesan untuk input x
input(x) // Masukkan nilai x
output("Masukkan bilangan n: ") // Tampilkan pesan untuk input n
input(n) // Masukkan nilai n

    found = false   // Inisialisasi variabel found

    // Loop untuk mencari digit x dalam bilangan n
    while n > 0 do
        digit = n mod 10
        if digit == x then
            found = true
            break
        endif
        n = n div 10
    endwhile

    // Output hasil pencarian
    if found then
        output("true")
    else
        output("false")
    endif

endprogram
