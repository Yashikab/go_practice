package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(64))
	fmt.Println("My favorite number is", r.Intn(10))
}
