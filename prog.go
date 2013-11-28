package main
import "fmt"
import "time"


func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
}

func main() {
	go ready("Tea", 2)			//2sec after print $W is ready
	go ready("Coffee", 1)		//1sec after print $W is ready
	fmt.Println("I'm waiting")	//0sec after print $W is ready
	time.Sleep(5 * time.Second) //timeout
	print(type(ready))
}



func conn(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
}




