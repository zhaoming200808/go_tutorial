package main

//import "os"
import "fmt"
import "reflect"

type T struct{
	a	int
	b	int
}

func (t T) getA() int{
	return t.a
}

func (t T) getB() int{
	return t.b
}

func main() {
	t := &T{
		a:10,
		b:20,
	}
	fmt.Printf("%#v\n",t)
	println("getA",t.getA())
	println("getB",t.getB())
	st := reflect.ValueOf(t)
	//mk := st.MapKeys()
	println(st.NumMethod())
}

func test_fun() {
	fmt.Printf("i am test func \n")
	return
}


