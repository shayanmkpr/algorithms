//go:build ignore

package examples

/*
================================================================
Algorithm: Merge Sort
================================================================

[EASY] LeetCode 88 - Merge Sorted Array
Merge nums2 into nums1 in-place. nums1 has length m+n with trailing zeros
to fit nums2.
Example: nums1=[1,2,3,0,0,0] m=3, nums2=[2,5,6] n=3 -> [1,2,2,3,5,6]
Hint: Fill from the back to avoid overwriting unread nums1 values.

[EASY] LeetCode 21 - Merge Two Sorted Lists
Merge two sorted linked lists into one sorted list.
Example: list1=[1,2,4], list2=[1,3,4]  ->  [1,1,2,3,4,4]
Hint: Use a dummy head and a tail pointer. Each step pick the smaller of
the two current heads; append the remaining tail at the end.

[EASY] LeetCode 1356 - Sort Integers by The Number of 1 Bits
Sort numbers by the number of set bits ascending; ties by value ascending.
Example: arr = [0,1,2,3,4,5,6,7,8]  ->  [0,1,2,4,8,3,5,6,7]
Hint: Stable sort using a custom comparator (popcount, then value). Merge
sort is naturally stable, ideal for "tie by original/secondary order".

[MEDIUM] LeetCode 148 - Sort List
Sort a singly linked list in O(n log n) time and O(1) extra space (besides
recursion stack).
Example: head = [4,2,1,3]  ->  [1,2,3,4]
Hint: Find the middle with slow/fast pointers, recursively sort each half,
then merge two sorted lists with a dummy node.

[HARD] LeetCode 23 - Merge k Sorted Lists
Given an array of k sorted linked lists, merge them into one sorted list.
Example: lists = [[1,4,5],[1,3,4],[2,6]]  ->  [1,1,2,3,4,4,5,6]
Hint: Repeatedly merge pairs of lists (divide & conquer). Total work
O(N log k) where N is the total number of nodes.
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 88: Merge Sorted Array -- O(m+n) time, O(1) space
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2) [EASY] LC 21: Merge Two Sorted Lists -- O(m+n)
func mergeTwoLists(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			cur.Next, a = a, a.Next
		} else {
			cur.Next, b = b, b.Next
		}
		cur = cur.Next
	}
	if a != nil {
		cur.Next = a
	} else {
		cur.Next = b
	}
	return dummy.Next
}

// 3) [EASY] LC 1356: Sort Integers by The Number of 1 Bits -- O(n log n)
func sortByBits(arr []int) []int {
	popcount := func(x int) int {
		c := 0
		for x > 0 {
			c += x & 1
			x >>= 1
		}
		return c
	}
	mergeSortBits(arr, 0, len(arr)-1, popcount)
	return arr
}

func mergeSortBits(a []int, lo, hi int, pc func(int) int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) / 2
	mergeSortBits(a, lo, mid, pc)
	mergeSortBits(a, mid+1, hi, pc)
	tmp := make([]int, hi-lo+1)
	i, j, k := lo, mid+1, 0
	for i <= mid && j <= hi {
		bi, bj := pc(a[i]), pc(a[j])
		if bi < bj || (bi == bj && a[i] <= a[j]) {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}
	for ; i <= mid; i, k = i+1, k+1 {
		tmp[k] = a[i]
	}
	for ; j <= hi; j, k = j+1, k+1 {
		tmp[k] = a[j]
	}
	copy(a[lo:hi+1], tmp)
}

// 4) [MEDIUM] LC 148: Sort List -- O(n log n)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return mergeTwo(sortList(head), sortList(mid))
}

func mergeTwo(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			cur.Next, a = a, a.Next
		} else {
			cur.Next, b = b, b.Next
		}
		cur = cur.Next
	}
	if a != nil {
		cur.Next = a
	} else {
		cur.Next = b
	}
	return dummy.Next
}

// 5) [HARD] LC 23: Merge k Sorted Lists -- O(N log k) via divide & conquer
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		merged := []*ListNode{}
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				merged = append(merged, mergeTwo(lists[i], lists[i+1]))
			} else {
				merged = append(merged, lists[i])
			}
		}
		lists = merged
	}
	return lists[0]
}
