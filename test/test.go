package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[0] += 1
	fmt.Println(m[1])
}
