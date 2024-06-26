package main

import "fmt"

func count(until int) {
	fmt.Println(until)
	if until == 0 {
		return
	}

	count(until - 1)
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
