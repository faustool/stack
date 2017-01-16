package stack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	stack := NewStack()

	value, err := stack.Pop()
	assert.NotNil(t, err)
	assert.Nil(t, value)

	value, err = stack.Peek()
	assert.NotNil(t, err)
	assert.Nil(t, value)

	stack.Push("element1")
	stack.Push("element2")

	value, err = stack.Peek()
	assert.Nil(t, err)
	assert.Equal(t, "element2", value)

	value, err = stack.Peek()
	assert.Nil(t, err)
	assert.Equal(t, "element2", value)

	value, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "element2", value)

	value, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "element1", value)

	value, err = stack.Pop()
	assert.NotNil(t, err)
	assert.Nil(t, value)
}