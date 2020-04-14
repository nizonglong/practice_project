package main

import "fmt"

func main() {
	const a int = 10
	const b, c = "12", 13
	const (
		d = 'd'
		e = "e"
		f = false
	)

	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d, "e=", e, "f=", f)
	fmt.Printf("d=%c , e=%s, f=%t\n", d, e, f)

	const (
		sunday = iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)

	fmt.Println(sunday + monday + tuesday + wednesday + thursday + friday + saturday)
}
