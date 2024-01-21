# 70. Palindrome Number (Easy)

LeetCode: https://leetcode.com/problems/palindrome-number/

Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:

```
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
```

Example 2:

```
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

Example 3:

```
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

# Solutions

The problem itself is not difficult to solve. We can solve the problem in two different ways; the first is pretty basic. We need to convert the input into a slice/array and compare the start with the end of the array. 

But the second one is more elaborated since we don't need the extra memory to copy the entry into an array; we just need to calculate the inverted number by guetting the remainder and the quotient of a division.

## 1. Slice solution

```golang
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
```

## 2. Math Solution

```golang
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original, reversed := x, 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return original == reversed
}
```

