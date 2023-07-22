package main

import "fmt"

var i, j int = 1, 2

func main() {
	// 初期化子(initializer)を与える場合は型を省略できる
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
