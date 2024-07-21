package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	flag := true
	if len(text) < 5 {
		flag = false
	}
	txtByte := []byte(text)

	for _, ch := range txtByte[0 : len(txtByte)-1] {
		if !((ch >= '1' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) {
			flag = false
			break
		}
	}

	if flag {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
