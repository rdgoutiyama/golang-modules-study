package main

import (
	"fmt"
	"br.com.rutiyama/mymodule"
)

func main() {
	message := greetings.Hello("Guys")
	fmt.Println(message)
}