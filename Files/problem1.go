package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	var sum int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		sum += n
	}
	writer := io.Writer(os.Stdout)
	writer.Write([]byte(strconv.Itoa(sum)))
}
