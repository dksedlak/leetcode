package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getSlice(head *ListNode) []int {
	result := []int{}

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

func isPalindrome(head *ListNode) bool {
	slice := getSlice(head)

	size := len(slice)

	for idx := 0; idx < size/2; idx++ {
		if slice[idx] != slice[size-1-idx] {
			return false
		}
	}

	return true
}
