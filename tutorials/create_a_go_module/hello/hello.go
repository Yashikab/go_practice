package main

import (
	"fmt"

	"github.com/Yashikab/go_practice/tutorials/create_a_go_module/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
