package linkedlist

import (
	"unsafe"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 借助map（时间复杂度O(N), 空间复杂度O(N)）
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pos := make(map[unsafe.Pointer]struct{})
	for headA != nil && headB != nil {
		if _, ok := pos[unsafe.Pointer(headA)]; ok {
			return headA
		} else {
			pos[unsafe.Pointer(headA)] = struct{}{}
		}
		if _, ok := pos[unsafe.Pointer(headB)]; ok {
			return headB
		} else {
			pos[unsafe.Pointer(headB)] = struct{}{}
		}
		headA = headA.Next
		headB = headB.Next
	}
	if headA == nil {
		for headB != nil {
			if _, ok := pos[unsafe.Pointer(headB)]; ok {
				return headB
			}
			headB = headB.Next
		}
	}
	if headB == nil {
		for headA != nil {
			if _, ok := pos[unsafe.Pointer(headA)]; ok {
				return headA
			}
			headA = headA.Next
		}
	}

	return nil
}

// 双循环(时间复杂度O(mn)，空间复杂度O(1))
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	for headA != nil {
		p := headB
		for p != nil {
			if p == headA {
				return p
			}
			p = p.Next
		}
		headA = headA.Next
	}

	return nil
}

// 双指针（时间复杂度O(N), 空间复杂度O(1)）
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
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


