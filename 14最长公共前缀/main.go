package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pre := strs[0] // flow
	for _, s := range strs {
		if s == "" {
			pre = ""
		}
		for j := 0; j < len(s) && j < len(pre); j++ {
			// flowht
			if s[j] != pre[j] {
				pre = s[:j]
				break
			}
			if j == len(s)-1 || j == len(pre) {
				pre = s
			}
		}
	}
	return pre
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"abab", "aba", ""}))
}
