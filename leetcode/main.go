package main

import (
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(twoEditWords([]string{"word", "note", "ants", "wood"}, []string{"wood", "joke", "moat"}))
	// fmt.Println(minSteps(5))
	// fmt.Println(addMinimum("abcabcababcc"))
	// fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	// fmt.Println(minimumSum(8, 5))
	// fmt.Println(longestContinuousSubstring("abcdacaba"))
	// fmt.Println(
	// 	countIslands(
	// 		[][]int{
	// 			{0, 2, 1, 0, 0},
	// 			{0, 5, 0, 0, 5},
	// 			{0, 0, 1, 0, 0},
	// 			{0, 1, 4, 7, 0},
	// 			{0, 2, 0, 0, 8},
	// 		},
	// 		5))
	// fmt.Println(numSquares(12))
	// fmt.Printf("%d \n", lengthOfLongestSubstring("tmmzuxt"))
	// fmt.Printf("%f \n", findMedianSortedArrays([]int{1,2}, []int{3, 4,5,6,9}))
	// fmt.Printf("%v", convert("PAYPALISHIRING", 4))
	// fmt.Printf("%v", convert("PA", 3))
	// fmt.Printf("%v", reverse(1234))
	// fmt.Printf("%v", myAtoi("words and 997"))
	// fmt.Printf("%v", isPalindrome(1234))
	// fmt.Println(isMatch("aaa", "aa*"))
	// fmt.Println(findMedianSortedArrays2([]int{}, []int{5, 6, 9}))
	// fmt.Println(findMedianSortedArrays2([]int{1, 2}, []int{3, 4}))
	// fmt.Println(maxArea([]int{1, 2, 3, 4, 5, 25, 24, 3, 4}))
	// fmt.Println(intToRoman(3749))
	// fmt.Println(romanToInt("MCMXCIV"))
	// fmt.Println(letterCombinations("23"))
	// fmt.Println(climbStairs(23))
	// fmt.Println(numTrees(23))
	// fmt.Println(isInterleave("afaf", "afaf", "afafafaf"))
	// fmt.Println(generateParenthesis(3))
	// fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	// fmt.Println(longestPalindrome("hellolll"))
	// fmt.Println("final", countSubstrings("ccc"))
	// fmt.Println("final", wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	// fmt.Println("***********************")
	// fmt.Println("final", wordBreak("catsanddog", []string{"cats", "dog", "sand", "and", "cat"}))
	// fmt.Println(missingNumber([]int{}))
	// fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	// fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	// fmt.Println(maxSubArray([]int{-2, -3, -5}))
	// fmt.Println(permute([]int{1, 2, 3}))
	// fmt.Println(sortColors([]int{2, 2, 2, 0, 2, 1, 1, 0}))
	// fmt.Println(longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}))
	// fmt.Println(numIslands([][]byte{
	// 	{'1', '1', '1', '1', '0'},
	// 	{'1', '1', '0', '1', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '0', '0', '0'},
	// }))
	// fmt.Println(isValidSudoku(
	// 	[][]byte{
	// 		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// 	},
	// ))
	// fmt.Println(spiralOrder(
	// 	[][]int{
	// 		{1, 2, 3, 4, 5},
	// 		{1, 1, 0, 1, 0},
	// 		{1, 1, 0, 0, 0},
	// 		{0, 0, 0, 0, 0},
	// 	}))
	// fmt.Println(canConstruct("aa", "baab"))
	// fmt.Println(lexicalOrder(12))
	// fmt.Println(firstUniqChar("helloh"))
}
