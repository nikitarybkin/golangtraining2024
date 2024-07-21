package main

import (
	"fmt"
	"strings"
)

// пакет используется для проверки ответа, не удаляйте его

type Battery struct {
	Power string
}

func (b Battery) String() string {
	return b.Power
}

func main() {
	batteryForTest := new(Battery)
	fmt.Scan(&batteryForTest.Power)
	pwr := strings.Count(batteryForTest.Power, "0")
	nStr := make([]rune, 0)
	nStr = append(nStr, '[')
	for i := 0; i < pwr; i++ {
		nStr = append(nStr, ' ')
	}
	for i := 0; i < 10-pwr; i++ {
		nStr = append(nStr, 'X')
	}
	nStr = append(nStr, ']')
	batteryForTest.Power = string(nStr)

}
