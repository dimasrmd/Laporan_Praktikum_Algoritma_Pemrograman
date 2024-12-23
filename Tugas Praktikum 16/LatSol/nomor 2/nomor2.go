package main

import "fmt"

func main() {
	var target, cek string
	var jumlah, poin_b, poin_c, n int
	fmt.Print("Masukkan kata yang dicari: ")
	fmt.Scan(&target)
	fmt.Print("Masukkan jumlah kata: ")
	fmt.Scan(&jumlah)
	n = 1
	if jumlah > 1 {
		for i := 1; i <= jumlah; i++ {
			fmt.Scan(&cek)
			if cek == target {
				if n == 1 {
					poin_b = i
					n++
				}
				poin_c += 1
			}
		}

		if poin_c > 0 {
			fmt.Printf("Apakah string (%s) ada dalam kumpulan %d data string tersebut? Ada.\n", target, jumlah)
		} else if poin_c == 0 {
			fmt.Printf("Apakah string (%s) ada dalam kumpulan %d data string tersebut? Tidak ada.\n", target, jumlah)
		}

		fmt.Printf("Pada posisi ke berapa string (%s) tersebut ditemukan? Posisi ke-%d\n", target, poin_b)
		fmt.Printf("Ada berapakah string (%s) dalam kumpulan %d data string tersebut? Ada %d\n", target, jumlah, poin_c)

		if poin_c >= 2 {
			fmt.Printf("Adakah sedikitnya dua string (%s) dalam kumpulan data? Ada\n", target)
		} else if poin_c < 2 {
			fmt.Printf("Adakah sedikitnya dua string (%s) dalam kumpulan data? Tidak ada\n", target)
		}
	} else {
		fmt.Print("Masukan kedua tidak valid.")
	}

}
