package main

import "fmt"

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	fmt.Println(nums1)
	fmt.Println(nums2)
	var merged []int
	var p1 int
	var p2 int
	p1, p2 = 0, 0
	cntr := 0

	if len(nums1) == 0 || len(nums1) == 0 {
		nums1 = nums2
	}
	for cntr < len(nums1)+len(nums2) && p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] < nums2[p2] {
			merged = append(merged, nums1[p1])
			fmt.Println("1 won")
			fmt.Println(merged)
			p1 += 1
			if p1 == len(nums1) {
				merged = append(merged, nums2[p2])
				fmt.Println("2 won")
				fmt.Println(merged)
				p2 += 1
			}

		} else {
			merged = append(merged, nums2[p2])
			fmt.Println("2 won")
			fmt.Println(merged)
			p2 += 1
			if p2 == len(nums2) {
				merged = append(merged, nums1[p1])
				fmt.Println("1 won")
				fmt.Println(merged)
				p1 += 1
			}

		}
		cntr += 1
		if (len(nums1)+len(nums2))%2 == 0 {
			if cntr == (len(nums1)+len(nums2))/2+1 {
				fmt.Println("up")
				fmt.Println(cntr)
				fmt.Println(merged[cntr-1])
				return float64(merged[cntr-2]+merged[cntr-1]) / 2
			}
		} else {
			if cntr == (len(nums1)+len(nums2))/2+1 {
				fmt.Println("down")
				fmt.Println(cntr)
				return float64(merged[cntr-1])
			}
		}
	}

	return 0
}
