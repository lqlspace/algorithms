package list

import (
	"testing"
	"strconv"
)

func TestPushSLList(t *testing.T) {
	list := New()
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	var str string
	for e := list.front; e != nil ; e = e.next {
		str += strconv.Itoa(e.Value.(int))
	}

	if str != "234" {
		t.Error("[error] push element error!")
		return
	}
	t.Log("push element succeed: ", str)
}

//测试反转功能
func TestReverseList(t *testing.T) {
	list := New()
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	list.ReverseList()

	var str string
	for e := list.front; e != nil; e = e.next {
		str += strconv.Itoa(e.Value.(int))
	}
	if str != "432" {
		t.Error("[error] reverse list error: ", str)
		return 
	}
	t.Log("reverse list succeed: ", str)
}
