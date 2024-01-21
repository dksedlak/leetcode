package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	number := strconv.Itoa(x)
	size := len(number)

	for idx := 0; idx < size/2; idx++ {
		if number[idx] != number[size-1-idx] {
			return false
		}
	}

	return true
}
