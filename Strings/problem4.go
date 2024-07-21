package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(text)

	var rsTwo []rune

	for i := 1; i < len(rs); i += 2 {
		rsTwo = append(rsTwo, rs[i])
	}

	fmt.Println(string(rsTwo))

}
