package main

// func merge(left, right *ListNode) *ListNode {
// 	dummy := &ListNode{}
// 	tail := dummy
// 	for left != nil && right != nil {
// 		if left.Val > right.Val {
// 			tail.Next = right
// 			right = right.Next
// 		} else {
// 			tail.Next = left
// 			left = left.Next
// 		}
// 		tail = tail.Next
// 	}
// 	if left != nil {
// 		tail.Next = left
// 	}
// 	if right != nil {
// 		tail.Next = right
// 	}
// 	return dummy.Next
// }

// func sortList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	slow, fast := head, head.Next
// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}
// 	middleNode := slow.Next
// 	slow.Next = nil
//
// 	left := sortList(head)
// 	right := sortList(middleNode)
//
// 	return merge(left, right)
// }
