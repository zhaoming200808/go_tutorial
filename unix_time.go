package main

//import "os"
import "fmt"
import "time"


func main() {
	n := 100
	//fmt.Printf(time.UTC,time.Local.String(),time.Now().Unix())
	fmt.Printf("%d\n",time.Now().UnixNano())
	time.Sleep( 10 * time.Nanosecond)
	fmt.Printf("%d\n",time.Now().UnixNano())
	time.Sleep( 10 * time.Nanosecond)
	fmt.Printf("%d\n",time.Now().UnixNano())
	time.Sleep( 10 * time.Nanosecond)
	fmt.Printf("%d\n",time.Now().UnixNano())
	time.Sleep( 10 * time.Nanosecond)
	fmt.Printf("%d\n",time.Now().UnixNano())
	print(n)
}

func test_fun() {
	fmt.Printf("i am test func \n")
	return
}


