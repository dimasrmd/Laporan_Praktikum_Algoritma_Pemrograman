package main
import (
	"fmt"
	"os"
	"bufio"
)
func main() {
	reader :=bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama anda: ")
	nama, _ :=reader.ReadString('\n')
	fmt.Print("Masukkan nim anda: ")
	nim, _ :=reader.ReadString('\n')
	fmt.Print("Masukkan kelas anda: ")
	kelas, _ :=reader.ReadString('\n')

	fmt.Println("\n---Resume Singkat---")
	fmt.Printf("Halo perkenalkan nama saya %s dengan nim %s, saya berasal dari kelas %s. ", nama, nim, kelas)
}