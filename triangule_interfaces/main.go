package main

import "fmt"

type square struct {
	sideLenght float64
}

type triangle struct {
	height float64
	base float64
}

type shape interface {
	getArea() float64
}

func main(){
	s := square{sideLenght: 4}
	printArea(s)
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func printArea (s shape) {
	fmt.Println(s.getArea())
}