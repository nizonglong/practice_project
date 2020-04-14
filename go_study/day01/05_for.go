package main

import "fmt"

func main() {
	s := "hello, world! 你好"

	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%c,", s[i])
	}

	fmt.Println()

	for _, r := range s { // rune
		fmt.Printf("%c,", r)
	}

	fmt.Println()

	n := len(s) - 1
	for n > 0 { // 类似while
		fmt.Printf("%c", s[n])
		n--
	}
}
