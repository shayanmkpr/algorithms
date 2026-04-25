package main

// list node was defined earlier somewhere
func hasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]bool)
	node := head
	nodeSeen := false
	for node != nil {
		nodeSeen = seen[node]
		if nodeSeen {
			return true
		}
		seen[node] = true
		node = node.Next
	}
	return false
}
