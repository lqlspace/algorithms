package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCQueue_AppendTail(t *testing.T) {
	q := ConstructorCQueue()
	res := q.DeleteHead()
	assert.Equal(t, res, -1)

	q.AppendTail(1)
	q.AppendTail(2)
	q.AppendTail(3)
	q.AppendTail(4)
	res = q.DeleteHead()
	assert.Equal(t, res, 1)

}
