package main

var memory map[int]int = map[int]int{}

func climbStairs(n int) int {
	if val, ok := memory[n]; ok {
		return val
	}

	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	first := climbStairs(n - 1)
	second := climbStairs(n - 2)
	memory[n] = first + second

	return memory[n]
}
