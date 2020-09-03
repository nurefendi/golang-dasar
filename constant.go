package main

import "fmt"

/*
* constant adalah variable yang nilainya tidak bisadiubah lagi setelah pertama diberi nilai
* cara pembuatan constant mirip dengan variable, yang membedakan hanya katakunci yang digunakan adalah const bukan var
* saat pembuatan constant, kita wajip langsung menginisialisasikan datanya
 */
func main() {
	const firstName = "fendi"
	const lastName = "abc"
	const value = 2000
	fmt.Println(firstName)

	const (
		namaPertama = "fendi"
		namakedua   = "abc"
	)
	fmt.Println(namaPertama)

}
