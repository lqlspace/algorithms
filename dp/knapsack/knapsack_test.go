package knapsack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnapsack(t *testing.T) {
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}
	n := 3 
	w := 4

	res := Knapsack(w, n, wt, val)

	assert.Equal(t, 6, res)
}
