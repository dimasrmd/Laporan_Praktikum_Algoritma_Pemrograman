package main
import "fmt"
func main() {
	var f float64
	fmt.Print("masukkan suhu dalam fahrenheit: ")
	fmt.Scan(&f)
	c:=(f-32)*5/9
	fmt.Printf("Jika di konverensikan ke celsius menjadi %.2f\n", c)
}