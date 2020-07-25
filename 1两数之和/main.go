package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	m := make(map[int]int)
	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			res[0] = v
			res[1] = i
		}
		m[num] = i
	}
	return res
}

func main() {
	nums := []int{7, 2, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
