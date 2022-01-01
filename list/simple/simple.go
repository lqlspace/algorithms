package simple

import (
	"fmt"
	"strings"
)

func CreateList(arr []int) *ListNode {
	sentinel := new(ListNode)
	last := sentinel

	for _, item := range arr {
		node := new(ListNode)
		node.Val = item
		last.Next = node
		last = last.Next
	}

	return sentinel.Next
}

func TraverseList(head *ListNode) {
	var builder strings.Builder
	for head != nil {
		builder.WriteString(fmt.Sprintf("%d ", head.Val))
		head = head.Next
	}

	println(builder.String())
}
