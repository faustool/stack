package stack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	arrayStack := ArrayStack{make([]string, 0)}

	arrayStack.Push("element1")
	arrayStack.Push("element2")

	value, error := arrayStack.Peak()
	assert.Nil(t, error)
	assert.Equal(t, "element2", value)

	value, error = arrayStack.Peak()
	assert.Nil(t, error)
	assert.Equal(t, "element2", value)

	value, error = arrayStack.Pop()
	assert.Nil(t, error)
	assert.Equal(t, "element2", value)

	value, error = arrayStack.Pop()
	assert.Nil(t, error)
	assert.Equal(t, "element1", value)

	value, error = arrayStack.Pop()
	assert.NotNil(t, error)
	assert.Equal(t, "", value)
}