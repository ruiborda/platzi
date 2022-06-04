package main

func Sum(x, y int) int {
	return x + y
}

func Fibonacci(n uint8) []uint64 {
	var fib []uint64
	a, b := uint64(0), uint64(1)
	for i := 0; i <= int(n); i++ {
		a, b = b, a+b
		fib = append(fib, a)
	}
	return fib
}

func FibonacciFunction(n uint8) uint64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FibonacciFunction(n-1) + FibonacciFunction(n-2)
}
