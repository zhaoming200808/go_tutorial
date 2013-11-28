package main

import (
	"sync"
)

var l sync.Mutex
var a string

func f() {
	a = "hello_world"
	print("21 ")
	l.Unlock()
	print("22 ")
}

func main() {
	l.Lock()
	print("1 ")
	f()
	print("2 ")
//	l.Lock()
	print("3 ")
//	l.Unlock()
	print(a)
}
// 1 21 22 2 3 hello_world
