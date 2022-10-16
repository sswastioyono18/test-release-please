package main

import "fmt"

func main() {
	fmt.Println("This is first commit")
	fmt.Println("123")
	fmt.Println("123")
	fmt.Println("123")
	fmt.Println("123")
	fmt.Println(add(1, 2))
}

func add(a, b int) int {
	if a == 1 {
		return 0
	}
	return a + b
}
