package main

func removeNthFromEnd0(head *ListNode, n int) *ListNode {
	cntr := 0

	var revList func(head *ListNode) *ListNode
	revList = func(head *ListNode) *ListNode {
		prev := head
		curr := head.Next
		var next *ListNode
		for curr != nil { // here we are just finding the end
			next = curr.Next // saving
			curr.Next = prev // reversing
			prev = curr
			curr = next // going for the next one
		}
		return prev
	}

	node := revList(head)
	for node.Next != nil && node.Next.Next != nil { // checking if the List has the array that I am looking for
		if cntr == n-1 {
			// rewire
			node.Next = node.Next.Next
		}
	}

	newHead := revList(node)
	return newHead
}
