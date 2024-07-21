package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	nums := strings.Split(s, ";")
	if len(nums) > 2 {
		panic("too many arguments")
	}
	num1, _ := Strtodigit(nums[0])
	num2, _ := Strtodigit(nums[1])
	fmt.Printf("%.4f", num1/num2)
}

func Strtodigit(s string) (float64, error) {
	var runes = make([]rune, 0)
	for _, ch := range s {
		if ch == ',' {
			runes = append(runes, rune('.'))
		} else if !unicode.IsSpace(ch) {
			runes = append(runes, rune(ch))
		}
	}
	return strconv.ParseFloat(string(runes), 64)
}
