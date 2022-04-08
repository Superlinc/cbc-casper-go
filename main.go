package main

import "fmt"

func main() {
	var bb t
	bb = &b{}
	bb.foo2()
}

type a int

func (receiver *a) foo() {
	fmt.Println("a")
}

func (receiver *a) foo2() {
	receiver.foo()
}

type b struct {
	*a
}

func (receiver *b) foo() {
	fmt.Println("b")
}

//func (receiver *b) foo2() {
//	receiver.foo()
//}

type t interface {
	foo()
	foo2()
}
