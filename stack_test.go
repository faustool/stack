package stack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	arrayStack := NewStack()

	value, err := arrayStack.Pop()
	assert.NotNil(t, err)
	assert.Nil(t, value)

	value, err = arrayStack.Peak()
	assert.NotNil(t, err)
	assert.Nil(t, value)

	arrayStack.Push("element1")
	arrayStack.Push("element2")

	value, err = arrayStack.Peak()
	assert.Nil(t, err)
	assert.Equal(t, "element2", value)

	value, err = arrayStack.Peak()
	assert.Nil(t, err)
	assert.Equal(t, "element2", value)

	value, err = arrayStack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "element2", value)

	value, err = arrayStack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "element1", value)

	value, err = arrayStack.Pop()
	assert.NotNil(t, err)
	assert.Nil(t, value)
}