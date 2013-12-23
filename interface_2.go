package main

import (
    "fmt"
//    "os"
)

// 定义 ISayHello 接口
type ISayHello interface {
	SayHello()
	Who()
}

//定义 Person 类
type Person struct {}
//实现 ISayHello 接口 必须实现所有接口
func (persion Person) SayHello() {
	fmt.Printf("Hello!\n")
}
func (persion Person) Who() {
	fmt.Printf("i am Person!\n")
}

//定义 Duck 类
type Duck struct {}
func (duck Duck) SayHello() {
	fmt.Printf("ga ga ga!\n")
}
func (duck Duck) Who() {
	fmt.Printf("i am Duck!\n")
}

//调用类的指定接口
func greeting(i ISayHello) {
	i.SayHello()
}

func call_who(i ISayHello) {
	i.Who()
}


//定义 Test 类
type Test struct {}

func main() {
	persion := Person{}
	duck := Duck{}
//	persion.SayHello()
//	duck.SayHello()
//	Test.print()
	greeting(persion)
	greeting(duck)
	call_who(persion)
	call_who(duck)
}

