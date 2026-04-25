package main

func reorderList(head *ListNode) {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil { // add the Next.Next check
		slow = slow.Next
		fast = fast.Next.Next
	}
	var prev *ListNode
	curr := slow.Next
	slow.Next = nil // breaking the second half from the first one
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	first := head
	second := prev

	for second != nil && first != nil {
		temp1, temp2 := first.Next, second.Next
		first.Next = second
		second.Next = temp1
		first, second = temp1, temp2
	}
}
