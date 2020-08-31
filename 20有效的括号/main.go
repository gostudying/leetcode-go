package main

import "fmt"
import "container/list"

func isValid(s string) bool {
	stack := list.New()
	m := map[int32]int32{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, v := range s {
		if stack.Len() == 0 {
			stack.PushBack(v)
		} else {
			// 判断当前字符是不是右括号，左括号直接压栈，右括号判断和栈顶元素的关系
			if m[v] == stack.Back().Value.(int32) {
				stack.Remove(stack.Back())
			} else {
				stack.PushBack(v)
			}
		}
	}
	return stack.Len() == 0
}

func main() {
	fmt.Println(isValid("()"))
}
