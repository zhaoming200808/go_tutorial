package main

import "fmt"
import "math"

func Add_1(x int, y int) int {
    return x + y
}

func add_2(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}


func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    fmt.Println(add_1(42, 13))
    fmt.Println(add_2(42, 13))
    a, b := swap("hello", "world")
    fmt.Println(a, b)
	fmt.Println(split(83))

	var x, y, z int
	var c, python, java bool
	fmt.Println(x, y, z, c, python, java)

    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
	// 函数也可以是值 （右值 )
	fmt.Println(hypot(3, 4))

    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
			"pos: ",pos(i),
			"neg: ",neg(-2*i),
        )
    }

    for i := 0; i < 10; i++ {
        fmt.Println(
			"pos: ",pos(i),
			"neg: ",neg(-2*i),
        )
    }



}





