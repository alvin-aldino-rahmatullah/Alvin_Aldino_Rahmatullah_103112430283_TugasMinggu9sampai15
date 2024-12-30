program Median
kamus
    j, y, bilangan, total : integer
algoritma
    input(y)
    total = 0
    for j = 1 to 9 do
        input(bilangan)
        total = total + bilangan
    endfor
    if total >= y * 5 then
        output("Median bernilai", y)
    else
        output("Median bernilai 0")
    endif
endprogram
