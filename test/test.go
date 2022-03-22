package main

import "fmt"

func main() {

	arr := make([]int, 2, 2)
	test(arr)
	fmt.Println(arr)

}

func test(arr []int) {
	arr = make([]int, 4, 4)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
}
