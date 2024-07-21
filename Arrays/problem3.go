package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	a = array[0]

	for i := range array {
		if array[i] > a {
			a = array[i]
		}
	}

	fmt.Println(a)

}
