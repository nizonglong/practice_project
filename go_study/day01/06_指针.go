package main

import (
	"fmt"
)

func main() {
	type data struct {
		a int
	}

	var d = data{1234}
	var p *data

	p = &d
	fmt.Printf("%p, %v\n", p, p.a)

}
