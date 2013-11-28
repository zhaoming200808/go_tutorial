package main

//import "os"
import "fmt"


// print 1 - 100
func main() {
//	fmt.Printf("Hello, world\n")
	s := fun_for(12690)

	fmt.Println(s)
	fmt.Println( fun_for_while(100) )

	fun_switch("test1")
	fun_switch("test2")
	fun_switch("test3")
	fun_switch("aaa")
}

func fun_for(max int) int {
	fmt.Printf("\n")
	sum := 0
	for i := 1 ; i <= max ; i++ {
		sum = sum + i
		if i%100 == 0 {
			fmt.Printf("%d\t",i)
		}
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
	return sum
}


func fun_switch(str string) int {
	switch str {
	case "test1":
		fmt.Printf("test1\n")
	case "test2","test3":
		fmt.Printf("testOhter: 2 or 3\n")
	case "2":
		fmt.Printf("int: 2\n")
	default:
		fmt.Printf("NoTest: %s\n",str)
	}
	return 0
}


func fun_for_while(num int) int {
//	i := 0
	sum := 0
	if i := 0 ; num > i {
		for i <= num {
			sum += i
			i ++ ;
		}
	}else {
		return -1
	}
	return sum
}




