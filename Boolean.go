package main

import "fmt"

func main() {
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)
	
	// mencari bilangan prima
	for i := 0; i < 10; i++ {
		counter := 0
		for j := 1; j <= i; j++ {
			if (i % j == 0){
				counter++
			}
		}
		
		if (counter == 2){
			fmt.Println(i, true)
		} else {
			fmt.Println(i, false)
		}
	}
}
