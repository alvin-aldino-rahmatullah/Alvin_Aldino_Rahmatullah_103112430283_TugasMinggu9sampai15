Program konsonan
Kamus:
    huruf : rune
Algoritma:
    output("Masukkan satu karakter: ")   // Tampilkan pesan untuk input karakter
    input(huruf)                         // Masukkan karakter

    if ('a' <= huruf AND huruf <= 'z') OR ('A' <= huruf AND huruf <= 'Z') then
        if huruf != 'a' AND huruf != 'e' AND huruf != 'i' AND huruf != 'o' AND huruf != 'u' AND
           huruf != 'A' AND huruf != 'E' AND huruf != 'I' AND huruf != 'O' AND huruf != 'U' then
            output("konsonan")           // Jika bukan vokal, karakter adalah konsonan
        else
            output("bukan konsonan")     // Jika vokal, bukan konsonan
        endif
    else
        output("bukan konsonan")         // Jika bukan huruf, bukan konsonan
    endif
EndProgram
