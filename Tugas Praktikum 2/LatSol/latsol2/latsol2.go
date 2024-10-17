package main
import "fmt"
func main() {
	//input
	var name, class, region, nim string
	fmt.Print("masukkan nama anda: ")
	fmt.Scan(&name)
	fmt.Print("Masukkan kelas anda: ")
	fmt.Scan(&class)
	fmt.Print("Masukkan NIM anda: ")
	fmt.Scan(&nim)
	fmt.Print("Masukkan asal daerah anda: ")
	fmt.Scan(&region)

	//output
	fmt.Printf("Halo semuanya, perkenalkan nama saya %s dengan NIM %s, saya dari kelas %s, lalu asal daerah saya dari %s, salam kenal semuanya.", name, nim, class, region)
}