package main

import "fmt"

/*
* Di go-Lang terkadang kita membutuhkan konversi
* hati - hati ketika konversi kebawahnya (int32 -> int8)
 */
func main() {

	var nilai32 int32 = 32344
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// konversi dari byte karakter dari string
	var name = "fendi"
	var e byte = name[0]
	var eString = string(e)

	fmt.Println(eString)

}
