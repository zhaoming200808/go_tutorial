package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) init_1(n float64) float64 {
	v.X = n
	v.Y = n
	// 在函数内部 打印值
	fmt.Println("at method init_1: ",v.X, v.Y)
	return n
}

func (v Vertex) init_2(n float64) float64 {
	v.X = n
	v.Y = n
	// 在函数内部 打印值
	fmt.Println("at method init_2: ",v.X, v.Y)
	return n
}

func main() {
	v := &Vertex{3, 4}
    v.init_1(1)
	fmt.Println(v.X, v.Y)

	v1 := Vertex{3, 4}
    v1.init_2(1)
	fmt.Println(v1.X, v1.Y)
}

