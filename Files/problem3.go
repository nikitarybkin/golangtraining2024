package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("E:\\Projects\\GoLang\\Study\\src\\Files\\task.data")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	reader := csv.NewReader(rd)
	reader.Comma = ';'
	rec, _ := reader.Read()
	for num, item := range rec {
		if item == "0" {
			fmt.Println(num + 1)
			break
		}
	}
}
