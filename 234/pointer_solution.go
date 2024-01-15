package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getMiddlePointer(head *ListNode) *ListNode {
	fastPointer, slowPointer := head, head
	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}
	return slowPointer
}

func reverseList(head *ListNode) *ListNode {
	reverse := head
	element := head.Next
	reverse.Next = nil

	for element != nil {
		tmp := element.Next
		element.Next = reverse
		reverse = element
		element = tmp
	}

	return reverse
}

func isPalindrome(head *ListNode) bool {
	middlePointer := getMiddlePointer(head)
	reverseList := reverseList(middlePointer)

	for reverseList != nil {
		if head.Val != reverseList.Val {
			return false
		}

		head = head.Next
		reverseList = reverseList.Next
	}

	return true
}
