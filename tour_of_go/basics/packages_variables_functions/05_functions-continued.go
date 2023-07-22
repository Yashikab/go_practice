package main

import "fmt"

// すべての引数が同じ型なら最後の型を残して省略できる
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
