# 70. Climbing Stairs (Easy)

LeetCode: https://leetcode.com/problems/climbing-stairs

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

```
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
```

Example 2:

```
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

# Solutions

The problem itself is not difficult to solve. We can easily find the recursive nature in the above problem. The person can reach $n^{th}$ nth stair from either $(n-1)^{th}$ stair or from $(n-2)^{th}$ stair. Hence, for each stair n, we try to find out the number of ways to reach $(n-1)^{th}$ stair and $(n-2)^{th}$ stair and add them to give the answer for the nth stair. Therefore the Recurrence relation for such an approach comes out to be: 

$climbStairs(n) = climbStairs(n-1) + climbStairs(n-2)$

The above expression is actually the expression for **Fibonacci numbers**, but there is one thing to notice, the value of ways(n) is equal to fibonacci(n+1). 

## 1. Recursion (Memoization - DP)

Basically, we just need to implement a method to calculate the  $n^{th}$ term and cached into a map.

```golang
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
```

## 2. Iterative

Find the $n^{th}$ term of the fibonnaci sequence.

```golang
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
```

