package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i // point to i
	// dereferencing
	fmt.Println(*p) // read i through the pointer
	// indirecting
	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j // point to j
	*p = *p / 37
	fmt.Println(j)
}
