package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// 関数内では := を使うと暗黙的な宣言ができる
	// 関数外ではできない
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
