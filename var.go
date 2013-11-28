package main

import "fmt"


func main() {

	var x, y, z int
	var c, python, java bool
	fmt.Println(x, y, z, c, python, java)

	var x1, y1, z1 int = 1, 2, 3
	var c1, python1, java1 = true, false, "no!"
	fmt.Println(x1, y1, z1, c1, python1, java1)

    var x2, y2, z2 int = 1, 2, 3
    c2, python2, java2 := true, false, "no!"

    fmt.Println(x2, y2, z2, c2, python2, java2)

}


