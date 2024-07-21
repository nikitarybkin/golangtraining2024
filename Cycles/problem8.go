package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, m int

	inCommon := ""
	fmt.Scan(&n, &m)

	if n == 0 && m == 0 {
		fmt.Print(0)
	} else {
		nStr, mStr := strconv.Itoa(n), strconv.Itoa(m)
		for nStr != "" {
			nDigit := nStr[0:1]
			mStrTmp := mStr
			for mStrTmp != "" {
				if mStrTmp[0:1] == nDigit {
					inCommon += nDigit + " "
					break
				} else {
					mStrTmp = mStrTmp[1:]
				}
			}
			nStr = nStr[1:]
		}
	}

	fmt.Print(inCommon)

}
