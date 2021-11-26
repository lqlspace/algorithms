package list

import (
	"unsafe"
)


// 哈希表(时间复杂度O（N），空间复杂度O(N))
func hasCycleMethod1(head *ListNode) bool {
	addrMap := make(map[unsafe.Pointer]struct{})
	for head != nil {
		if _, ok := addrMap[unsafe.Pointer(head)]; ok {
			return true
		}
		addrMap[unsafe.Pointer(head)] = struct{}{}
		head = head.Next
	}

	return false
}

// 快慢指针（时间复杂度O(N)，空间复杂度O(1)）
func hasCycleMethod2(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil {
		if fast == slow {
			return true
		}

		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}
