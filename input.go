package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("G'day %s!", name)
	fmt.Println()
}
