package main

import "fmt"

func main() {
	// familiar syntax
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	// idiomatic loop syntax
	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// while simulator
	// lesson said to put i := 0, got error abount new variable
	// not being initialized
	i = 0
	for {
		if i == 10 {
			break
		}
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// This is a slice but range also works with arrays
	aSlice := []int{-1, -2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index:", i, "value:", v)
	}
}
