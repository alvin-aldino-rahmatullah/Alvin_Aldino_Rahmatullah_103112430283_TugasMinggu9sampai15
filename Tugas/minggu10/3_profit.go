package main
import "fmt"

func main(){
   var bulan1, bulan2 float64

   fmt.Print("Masukkan pendapatan bulan pertama : ")
   fmt.Scan(&bulan1)
   fmt.Print("Masukkan pendapatan bulan kedua : ")
   fmt.Scan(&bulan2)

   if bulan2 > bulan1 {
	  total := bulan2 - bulan1 
	  fmt.Print("Pendapatan meningkat sejumlah ", total)
   } else if bulan1 > bulan2 {
	 total := bulan1 - bulan2 
	 fmt.Print("Penurunan sebesar ", total)
   } else if bulan1 == bulan2 {
	 fmt.Print("tetap")
   }
}