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
	for i := 0; i < len(rs); i++ {
		rsTwo = append(rsTwo, rs[len(rs)-1-i])
	}

	fmt.Println(string(rsTwo))
	fmt.Println(text)
	if text == string(rsTwo) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}

}
