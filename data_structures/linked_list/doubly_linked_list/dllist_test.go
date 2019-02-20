package dllist

import (
	"testing"
)

func TestPush(t *testing.T) {
	l := New()
	l.PushFront("a")
	l.PushFront("b")
	l.PushFront("c")

	if l.Len() != 3 {
		t.Error("[error] PushFront error, length is wrong !")
	}

	var str string
	for e := l.Front(); e != nil; e = e.Next() {
		str += e.Value.(string)
	}
	t.Log(str)
	if str != "cba" {
		t.Error("[error] Element not right!")
	}

	str = ""
	for e := l.Back(); e != nil; e = e.Prev() {
		str += e.Value.(string)
	}
	if str != "abc" {
		t.Error("[error] Element not right!")
	}
	t.Log("Push element ok: ", str)
}
