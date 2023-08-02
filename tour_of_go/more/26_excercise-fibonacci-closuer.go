package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib0 := 0
	fib1 := 1
	fib2 := fib0 + fib1
	return func() int {
		ans := fib0
		fib0 = fib1
		fib1 = fib2
		fib2 = fib0 + fib1
		return ans
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
