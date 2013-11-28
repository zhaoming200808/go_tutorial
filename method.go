package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}


func main() {
method_2_msg := `
可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
不能对来自其他包的类型或基础类型定义方法。
`
method_3_msg := `
一个是在 *Vertex 指针类型上，而另一个在 MyFloat 值类型上。 
有两个原因需要使用指针接收者。
首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。
其次，方法可以修改接收者指向的值。
`
    v := &Vertex{3, 4}
    fmt.Println(v.X,v.Y)
    fmt.Println(v.Abs())

    fmt.Println(method_2_msg)

    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())


    fmt.Println(method_3_msg)
	v = &Vertex{3, 4}
    v.Scale(5)
    fmt.Println(v.X)
    fmt.Println(v.Y)
    fmt.Println(v, v.Abs())



}



