package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	aSlice := make([]int, N)
	for i := range aSlice {
		fmt.Scan(&aSlice[i])
	}

	for i := 0; i < N; i += 2 {
		fmt.Print(aSlice[i], ` `)
	}

}
