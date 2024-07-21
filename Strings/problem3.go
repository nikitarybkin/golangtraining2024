package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	X, _ := reader.ReadString('\n')
	S, _ := reader.ReadString('\n')
	fmt.Println(X)
	fmt.Println(S)

	fmt.Println(strings.Index(X, S))

}
