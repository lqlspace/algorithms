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

func TestRemove(t *testing.T) {
	list := New()
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	
	list.Remove(3)

	var str string
	for e := list.front; e != nil; e = e.next {
		str += strconv.Itoa(e.Value.(int))
	}
	if str != "24" {
		t.Error("[error] remove value 3 error: ", str)
		return
	}

	t.Log("remove value 3 secceed: ", str)
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

//测试递归反转功能
func TestReverseListRecursive(t *testing.T) {
	list := New()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	list.ReverseListRecursive()

	var str string
	for e := list.front; e != nil; e = e.next {
		str += strconv.Itoa(e.Value.(int))
	}
	if str != "4321" {
		t.Error("[error] reverse list recursively error: ", str)
		return
	}

	t.Log("reverse list recursively succeed: ", str)
}
