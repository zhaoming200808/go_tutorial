package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

var (
    p1 = Vertex{1, 2}  // 类型为 Vertex
    q1 = &Vertex{1, 2} // 类型为 *Vertex
    r1 = Vertex{X: 1}  // Y:0 被省略
    s1 = Vertex{}      // X:0 和 Y:0
)

func main() {
    fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{X : 1,Y : 2})

	v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)

    p := Vertex{1, 2}
    q := &p
    q.X = 1e9
    fmt.Println(p)

	fmt.Println(p, q)
	fmt.Println(p1, q1, r1, s1)

    fmt.Println("--------point---------")
	// var t *T = new(T)
	// t := new(T)
	v2 := new(Vertex)
    fmt.Println(v2)
    v2.X, v2.Y = 11, 9
    fmt.Println(v2)



}





