package random

import (
	"testing"
)

func TestRandomIntSlice(t *testing.T) {
	arr := RandomIntSlice(10, 30, 35)
	t.Log(arr)

	arr = RandomDistinctIntSlice(2, 1, 3)
	t.Log(arr)
}
