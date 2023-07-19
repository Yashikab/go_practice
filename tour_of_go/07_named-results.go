package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// naked return と呼ぶ
	// 短い関数でのみ利用すべき
	return
}

func main() {
	fmt.Println(split(17))
}
