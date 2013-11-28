package main

//import "os"
import "fmt"

func main() {
	fmt.Printf("Hello, world\n")
	const (
		a = iota
		b
		c
		d
		e
		f
	)
	i1 := 1
	i2 := 2
	i3 := 3

	fmt.Printf("%d\n",a)
	fmt.Printf("%d\n",b)
	fmt.Printf("%d\n",c)
	fmt.Printf("%d\n",d)
	fmt.Printf("%d\n",e)
	fmt.Printf("%d\n",f)
}

func test_fun() int {
	fmt.Printf("i am test func \n")
	return 0
}


