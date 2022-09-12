package main

import "fmt"

func main() {
	fmt.Println("This is first commit")
	fmt.Println(add(1, 2))
}

func add(a, b int) int {
	if a == 1 {
		return 0
	}

	if a == 2 {
		return 2
	}
	return a + b
}
