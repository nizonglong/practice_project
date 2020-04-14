package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println("a[1]=", a[1])

	b := make([]int, 3)
	b[1] = 2
	fmt.Println("b[1]=", b[1])

}
