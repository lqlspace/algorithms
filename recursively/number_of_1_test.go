package recursive

import (
	"testing"
)

func TestNumberOf1(t *testing.T) {
	n := NumberOf1(12)
	if n != 5 {
		t.Error("[error] error number: ", n)
	}
	t.Log("succeed: ", n)
}
