package main

import "fmt"
import "time"

func main() {
	print("start\n")

	t0 := time.Now()
	time.Sleep(1000 * time.Millisecond)
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	fmt.Printf("local data: %v\n",time.Now())
	print("end\n")
}

func test_fun() int {
	fmt.Printf("i am test func \n")
	return 0
}


