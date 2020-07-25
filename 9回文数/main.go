package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	var sb strings.Builder
	for i := len(xStr) - 1; i >= 0; i-- {
		sb.WriteByte(xStr[i])
	}
	if sb.String() == xStr {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(12))
}
