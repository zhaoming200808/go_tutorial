package main

//import "os"
import "fmt"
func main() {
	n := 100
	println(n)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Runtime error caught: %v\n", r)
		}
	}()

	test()
	test_panic("404")
	n=n+100
	println(n)

	panic("3")
	n=n+100
	println(n)
}

func test_panic(x string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Runtime error caught: %v\n", r)
		}
	}()
	panic(x)
}

func test(){
	test_panic("abc")
}


func test_fun() int {
	fmt.Printf("i am test func \n")
	return 0
}


