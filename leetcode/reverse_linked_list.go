package main

import "fmt"

func reverseLinkedList() {
	type Node struct {
		val  int
		next *Node
	}
	var appendNode func(head *Node, n int)
	appendNode = func(head *Node, n int) {
		// getting to the end of the list
		node := head
		for node.next != nil {
			node = node.next
		}
		for i := 1; i <= n; i++ {
			newNode := &Node{ // to make sure we are adding to the heap not to the stack
				val:  i,
				next: nil,
			}
			node.next = newNode
			node = newNode
		}
	}
	head := &Node{
		val:  0,
		next: nil,
	}
	var printList func(head *Node)
	printList = func(head *Node) {
		node := head
		for node != nil {
			fmt.Println(node.val)
			node = node.next
		}
	}
	// reverse list in place (updates head via **Node)
	var rev func(head *Node) *Node
	rev = func(head *Node) *Node {
		var prev *Node
		curr := head
		for curr != nil {
			next := curr.next
			curr.next = prev
			prev = curr
			curr = next
		}
		return prev
	}
	// create list
	appendNode(head, 5)
	fmt.Print("Original: ")
	printList(head)
	head = rev(head)
	fmt.Print("Reversed: ")
	printList(head)
}
