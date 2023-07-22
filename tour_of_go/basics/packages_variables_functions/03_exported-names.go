package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(math.pi) // 小文字は呼び出せずエラーになる
	fmt.Println(math.Pi) // 大文字は exported namesなので外部から呼び出せる
}
