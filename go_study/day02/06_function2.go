package main

import "fmt"

// 变参
func funcTest(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}

func main() {
	funcTest(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
}
