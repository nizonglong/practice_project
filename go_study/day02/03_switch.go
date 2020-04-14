package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	i := 2
	switch i {
	case x[1]:
		fmt.Println("a")
	case 1, 3:
		fmt.Println("b")
	default:
	}
	fmt.Println("c")

	// 如需要继续下⼀一分⽀支，可使⽤用 fallthrough，但不再判断条件
	x2 := 10
	switch x2 {
	case 10:
		fmt.Println("a")
		fallthrough
	case 0:
		fmt.Println("b")
	}

	// 省略条件表达式，可当 if...else if...else 使⽤用
	switch {
	case x[1] > 0:
		fmt.Println("a")
	case x[1] < 0:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
	switch i := x[2]; { // 带初始化语句
	case i > 0:
		fmt.Println("a")
	case i < 0:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}
