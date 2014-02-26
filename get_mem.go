package main

import "time"

func main() {
	var arr [10000]int
	for i:=1 ; i < 10000 ; i++ {
		arr[i] = i
	}
	println("ok")
	time.Sleep(100 * time.Second)
}


