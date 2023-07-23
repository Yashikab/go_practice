package main

import "fmt"

func main() {
	// 呼び出し元の関数の終わりまで実行を遅延させる
	// 今回はmain関数
	// 渡した関数の引数はすぐに評価されるが、関数自体はmainがreturnするまで実行されない
	defer fmt.Println("world")
	fmt.Println("hello")
}
