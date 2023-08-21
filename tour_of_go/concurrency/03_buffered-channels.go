package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	go func() {
		ch <- 3
		ch <- 4
	}()
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 5
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
