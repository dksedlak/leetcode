package main

func climbStairs(n int) int {
	prev := 0
	current := 1
	total := 0

	for i := 0; i < n; i++ {
		total = prev + current
		prev = current
		current = total
	}

	return total
}
