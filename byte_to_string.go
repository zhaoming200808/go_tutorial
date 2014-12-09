package main

//import "os"
import "fmt"

func main() {
	n := 100
	println(n)
	test_fun()
}

func test_fun() {
	byte_1 := make([]byte, 4096)
	byte_1 = []byte("i am test")
	str_1 := "i am test"

	str_str := str_1
	byte_str := string(byte_1)

	fmt.Printf("byte_str |%v|\n", byte_str)
	fmt.Printf("str_str  |%v|\n", str_str)
	fmt.Printf("i am test func \n")
	fmt.Printf("%d %d\n", len(str_str), len(byte_str))
	if byte_str == str_str {
		fmt.Printf("byte_str eq str_str \n")
	} else {
		fmt.Printf("byte_str ne str_str \n")
	}
	return
}
