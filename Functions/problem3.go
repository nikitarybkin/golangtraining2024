package main

import "fmt"

func main() {
	fmt.Println(vote(0, 0, 1))

}

func vote(x int, y int, z int) int {
	zeroC := 0
	if x == 0 {
		zeroC++
	}
	if y == 0 {
		zeroC++
	}
	if y == 0 {
		zeroC++
	}
	if zeroC >= 2 {
		return 0
	} else {
		return 1
	}
}
