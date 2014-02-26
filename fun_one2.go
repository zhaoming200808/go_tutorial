package main

import "fmt"
import "time"
import "sync"

type Host struct {
	x int
}
func NewHost() *Host {
	return &Host{
		x:0,
	}
}

func (h *Host) t1(n,t int) {
	h.x = n
	println("i am ",n)
	time.Sleep(time.Duration(t) * time.Second)
}


func (h *Host) t2() {
	h.x = 10
	println("i am ",h.x)
	time.Sleep(time.Duration(5) * time.Second)
}

func (h *Host) t3() {
	h.x = 10
	println("i am ",h.x)
	time.Sleep(time.Duration(5) * time.Second)
}

var once sync.Once
func main() {
	host := NewHost()
	for i:=1 ; i<= 10 ; i++ {
		var once sync.Once
		//host.t1(i,5)
		go func() {
			once.Do(host.t2)
		}()
	}
	time.Sleep(6 * time.Second)
	fmt.Printf("x is: |%d|\n",host.x)
}



