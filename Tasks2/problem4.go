package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var x int

	for _, ch := range text {
		x = int(ch-'0') * int(ch-'0')
		fmt.Print(x)
	}
	strconv.ParseUint(strconv.Itoa(x), 10, 64)
}
