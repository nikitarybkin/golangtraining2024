package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(text)

	var rsTwo []rune

	for _, letter := range rs {
		if strings.Count(text, string(letter)) == 1 {
			rsTwo = append(rsTwo, letter)
		}
	}

	fmt.Println(string(rsTwo))

}
