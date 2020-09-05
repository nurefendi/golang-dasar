package main

import "fmt"

/*
* Type Declaration adalah kemampuan membuat ulang tipe baru dari tipe data yang sudah ada
* Biasanya digunakan untuk membuat alias terhadap data yang sudah ada, dengan tujuan agar lebih mudah dimengerti
 */

func main() {
	type noKTP string
	type Merried bool

	var ktpfendi noKTP = "1233234"
	var isMerried Merried = true
	fmt.Println(ktpfendi)
	fmt.Println(isMerried)
}
