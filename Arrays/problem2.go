package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	sl := make([]int, N)

	for i := range sl {
		fmt.Scan(&sl[i])
	}

	fmt.Print(sl[3])

}
