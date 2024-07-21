package main

import "fmt"

func main() {
	var N, counter int
	fmt.Scan(&N)
	counter = 0

	slice := make([]int, N)

	for i := range slice {
		fmt.Scan(&slice[i])
		if slice[i] > 0 {
			counter++
		}
	}

	fmt.Println(counter)

}
