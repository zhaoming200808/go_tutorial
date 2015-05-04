package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)
	routines := 10

	for i := 0; i < routines; i++ {
		go func(ch chan int, i int) {
			<-ch
			fmt.Printf("go routine %d recv exit signal, exiting.\n", i)
		}(ch, i)
	}

	close(ch)

	time.Sleep(2 * time.Second)
}
