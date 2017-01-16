package stack

import (
	"errors"
)

type Stack interface {
	Push(value interface{})
	Pop() (interface{}, error)
	Peak() (interface{}, error)
}

type ArrayStack struct {
	data [] interface{}
}

func NewStack() Stack {
	return &ArrayStack{make([]interface{}, 0)}
}

func (stack *ArrayStack) Push(value interface{}) {
	stack.data = append(stack.data, value)
}

func (stack *ArrayStack) Pop() (interface{}, error) {
	length := len(stack.data)
	if (length == 0) {
		return nil, errors.New("Cannot Pop from an empty Stack")
	} else {
		lastIndex := length - 1
		last := stack.data[lastIndex]
		stack.data = stack.data[:lastIndex]
		return last, nil
	}
}

func (stack ArrayStack) Peak() (interface{}, error) {
	length := len(stack.data)
	if (length == 0) {
		return nil, errors.New("Cannot Peak from an empty Stack")
	} else {
		return stack.data[length -1], nil
	}
}