package stack

import "fmt"

type Stack interface {
	Push(value string)
	Pop() (string, error)
	Peak() (string, error)
}

type ArrayStack struct {
	data []string
}

func (stack *ArrayStack) Push(value string) {
	stack.data = append(stack.data, value)
}

func (stack *ArrayStack) Pop() (string, error) {
	len := len(stack.data)
	if (len == 0) {
		return "", fmt.Errorf("Cannot Pop from an empty Stack")
	} else {
		lastIndex := len - 1
		last := stack.data[lastIndex]
		stack.data = stack.data[:lastIndex]
		return last, nil
	}
}

func (stack ArrayStack) Peak() (string, error) {
	len := len(stack.data)
	if (len == 0) {
		return "", fmt.Errorf("Cannot Peak from an empty Stack")
	} else {
		return stack.data[len -1], nil
	}
}