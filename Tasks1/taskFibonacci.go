package main

import "fmt"

func main() {
	var counter, A int
	fmt.Scan(&A)
	counter = 2
	fibNumber, fibNumberPrev := 1, 1
	result := -1

	for fibNumber <= A {
		if fibNumber == A {
			result = counter
		}
		fibNumberPrev, fibNumber = fibNumber, fibNumberPrev+fibNumber
		counter++
	}

	fmt.Print(result)

}
