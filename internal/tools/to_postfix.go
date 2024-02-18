package tools

import "strings"

func ToPostfix(s string) string {
	var stack Stack
	postfix := ""
	length := len(s)

	for i := 0; i < length; i++ {
		char := string(s[i])

		if char == " " {
			continue
		}

		if char == "(" {
			stack.Push(char)
		} else if char == ")" {
			for !stack.Empty() {
				str, _ := stack.Top().(string)
				if str == "(" {
					break
				}
				postfix += " " + str
				stack.Pop()
			}
			stack.Pop()
		} else if !IsOperator(s[i]) {
			j := i
			number := ""

			for ; j < length && IsOperand(s[j]); j++ {
				number += string(s[j])
			}
			postfix += " " + number
			i = j - 1
		} else {
			for !stack.Empty() {
				top, _ := stack.Top().(string)
				if top == "(" || !HasHigherPrecedence(top, char) {
					break
				}
				postfix += " " + top
				stack.Pop()
			}
			stack.Push(char)
		}
	}
	for !stack.Empty() {
		str, _ := stack.Pop().(string)
		postfix += " " + str
	}
	return strings.TrimSpace(postfix)
}
