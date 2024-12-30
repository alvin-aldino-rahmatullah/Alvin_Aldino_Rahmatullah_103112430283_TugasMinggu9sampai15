Program tigadanlima
Kamus:
    n : integer
Algoritma:
    output("Masukkan bilangan bulat: ")  // Tampilkan pesan untuk input bilangan
    input(n)                             // Masukkan bilangan bulat

    if n mod 3 == 0 AND n mod 5 == 0 then
        output("kelipatan 3")            // Jika bilangan kelipatan 3
        output("kelipatan 5")            // Jika bilangan kelipatan 5
    else if n mod 5 == 0 then
        output("kelipatan 5")            // Jika hanya kelipatan 5
    else if n mod 3 == 0 then
        output("kelipatan 3")            // Jika hanya kelipatan 3
    else
        output("")                       // Tidak menampilkan apapun
    endif
EndProgram
