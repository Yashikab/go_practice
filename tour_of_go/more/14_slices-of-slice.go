package main

import (
	"fmt"
	"string"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]stirng{"_", "_", "_"},
		[]stirng{"_", "_", "_"},
		[]stirng{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

}
