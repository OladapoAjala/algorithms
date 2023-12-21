func Fibonacci(n int, memo map[int]int) int {
	if n <= 2 {
		return 1
	}
	if f, ok := memo[n]; ok {
		return f
	}

	f := Fibonacci(n-1, memo) + Fibonacci(n-2, memo)
	memo[n] = f
	return f
}

