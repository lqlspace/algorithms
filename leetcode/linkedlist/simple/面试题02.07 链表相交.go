package simple

import (
	"unsafe"
)

// 借助hash表
func getIntersectionNodeM1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	m := make(map[unsafe.Pointer]struct{})
	for headA != nil {
		m[unsafe.Pointer(headA)] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := m[unsafe.Pointer(headB)]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

// 双指针
func getIntersectionNodeM2(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}

	return p1
}
