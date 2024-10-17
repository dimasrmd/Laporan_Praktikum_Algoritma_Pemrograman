package main

import "fmt"

func main() {
    var (
    satu, dua, tiga, temp string
    )
    fmt.Print("Masukan input 3 kata: ")
    fmt.Scanln(&satu, &dua, &tiga)
    
    fmt.Println("Output awal = " + satu + " " +
    dua + " " + tiga)
    temp = satu
    satu = dua
    dua = tiga
    tiga = temp
    fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
