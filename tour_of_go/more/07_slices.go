package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// 最後の要素は除いた区間になる
	var s []int = primes[1:4]
	fmt.Println(s)
}
