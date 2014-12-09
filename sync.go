package main

import (
    "sync"
	"time"
)

var l sync.Mutex
var a string

func f() {
	l.Lock()
    defer l.Unlock()
	println("sleep 3 second")
	time.Sleep(3 * time.Second)
}

func main() {
	go f()
	go f()

	time.Sleep(10 * time.Second)
}

