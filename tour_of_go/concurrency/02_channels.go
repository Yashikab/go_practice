package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0

	time.Sleep(time.Duration(s[2]*100) * time.Millisecond)
	for _, v := range s {
		sum += v
	}
	c <- sum //send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)
	// x, y := <-c, <-c //receive from c
	// fmt.Println(x, y, x+y)
	x := <-c
	y := <-c
	fmt.Println(x, y)
}
