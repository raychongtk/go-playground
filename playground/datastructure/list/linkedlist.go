package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

// https://leetcode.com/problems/merge-in-between-linked-lists/
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var step int
	n1, n2 := &ListNode{}, &ListNode{}
	head, head2 := list1, list2

	for head != nil {
		if step+1 == a {
			n1 = head
		}
		if step == b {
			n2 = head.Next
			break
		}
		head = head.Next
		step++
	}

	n1.Next = list2
	for head2.Next != nil {
		head2 = head2.Next
	}
	head2.Next = n2

	return list1
}
