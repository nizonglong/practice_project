package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		} else {
			defer fmt.Printf("%d ", i)
		}
	}
}
