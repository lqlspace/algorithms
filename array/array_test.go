package array

import (
	"testing"
)

func TestFind(t *testing.T) {
	arr := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	found := Find(arr, 4, 4, 7)
	t.Log(found)

	found = Find(arr,4, 4, 5)
	t.Log(found)
}

func TestFind2(t *testing.T) {
	str  := "1234"

	t.Log(len(str))
}
