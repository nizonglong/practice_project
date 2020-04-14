package main

import "fmt"

func main() {

	s := "abc"
	for i := range s { // 忽略 2nd value，支持 string/array/slice/map
		fmt.Printf("%c", s[i])
	}
	fmt.Println()
	for _, c := range s { // 忽略 index
		fmt.Printf("%c", c)
	}
	fmt.Println()
	for range s { // 忽略全部返回值，仅迭代

	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m { // 返回 (key, value)
		fmt.Println(k, v)
	}
}
