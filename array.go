package main

import "fmt"

func main()  {
	var names [3]string

	names[0] = "No 1"
	names[1] = "No 2"
	names[2] = "No 3"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])



	var  value = [3] int {
		1,2,3,
	}

	fmt.Println(value)
	fmt.Println(value[0])
	fmt.Println(value[1])
	fmt.Println(value[2])


	var totaldata = len(names)
	fmt.Println(totaldata)
	println(len(value))

}