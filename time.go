package main

import "fmt"
import "time"

func main() {
	print("start\n")

	//the_time, err := time.Parse("2006-01-02 15:04:05", "[30/Dec/2014:03:23:05 +0800]")
	the_time, err := time.Parse("[02/Jan/2006:15:04:05 +0800]", "[30/Dec/2014:03:23:05 +0800]")
	if err != nil {
		print(err)
	}
	//println(the_time)
	fmt.Printf("The Nginx Time is %v\n", the_time.Format("2006-01-02 15:04:05"))
	println(time.Now().Unix())
	t0 := time.Now()
	time.Sleep(1000 * time.Millisecond)
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	println(time.Now().Unix())
	fmt.Printf("local data: %v\n", time.Now())
	print("end\n")
}

func test_fun() int {
	fmt.Printf("i am test func \n")
	return 0
}
