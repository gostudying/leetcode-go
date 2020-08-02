package main

import "fmt"

// 从右到左遍历罗马数字的字符，将罗马字符映射为对应的阿拉伯数字，若当前的数字大于或等于前一个数字，则加，否则减。
func romanToInt(s string) int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res += m[s[len(s)-1]]
	for i := len(s) - 1; i > 0; i-- {
		if m[s[i-1]] >= m[s[i]] {
			res += m[s[i-1]]
		} else {
			res -= m[s[i-1]]
		}
	}
	return res
}

// 解法1
func romanToInt1(s string) int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			res += m[s[i]]
			break
		}
		if s[i] == 'I' {
			if s[i+1] == 'V' {
				res += 4
				i++
			} else if s[i+1] == 'X' {
				res += 9
				i++
			} else {
				res += m[s[i]]
			}
		} else if s[i] == 'X' {
			if s[i+1] == 'L' {
				res += 40
				i++
			} else if s[i+1] == 'C' {
				res += 90
				i++
			} else {
				res += m[s[i]]
			}
		} else if s[i] == 'C' {
			if s[i+1] == 'D' {
				res += 400
				i++
			} else if s[i+1] == 'M' {
				res += 900
				i++
			} else {
				res += m[s[i]]
			}
		} else {
			res += m[s[i]]
		}

	}
	return res
}

func main() {
	fmt.Println(romanToInt("LVIII"))
}
