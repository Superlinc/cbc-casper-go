package main

import "fmt"

func main() {
	a := make([][]int, 0, 4)
	test(a)
}

func test(a interface{}) {
	for _, v := range a.([][]int) {
		fmt.Println(v)
	}
}
