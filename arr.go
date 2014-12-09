package main

//import "os"
import "fmt"
import "strconv"

type S struct {
}



func main() {
	fmt.Printf("Hello, world\n")
	// arr: 0 - 0
	var a1 [10]int
	// arr:1 - 10 
	a2 := [10]int{1,2,3,4,5,6,7,8,9,10}
	// 1 2 0 0 0 6 0 0 0 0
	a3 := [10]int{1,2,5:6}
	// 二维数组 两个包含4个int元素的数组
	// arr[1] 1 2 3 0
	// arr[2] 5 6 7 8
	a4 := [2][4]int{[4]int{1,2,3}, [4]int{5,6,7,8} }
	//6个元素 分别为 1 2 3 0 0 6
	a5 := [...]int{1,2,3,5:6}

	fmt.Printf("---------array------------\n")
	fmt.Println("a1: ",a1)
	fmt.Println("a2: ",a2)
	fmt.Println("a3: ",a3)
	fmt.Println("a4: ",a4)
	fmt.Println("a5: ",a5)

	fmt.Printf("---------slice------------\n")
	// slice是一个指针，它指向的是一个array机构，它有两个基本函数len和cap。
	// slice 3 4 5 
	// slice 自动增长 不适合大数据 append 操作 
	sl := a2[2:5]
	sl1 := make([]int, 10)
	sl2:=[]string{}
	sl2 = append(sl2, "hello","word")

	fmt.Println("s1: ",sl)
	sl = append(sl,1,2,3,4,5,6,7)
	fmt.Println("s1: ",sl)
	fmt.Println("sl: ",sl1)
	fmt.Println("sl2: ",sl2)

	fmt.Printf("----------other------------\n")
	//len 元素个数 cap实际大小
	fmt.Printf("array len: %d\n",len(a1))
	fmt.Printf("array cap: %d\n",cap(a1))
	fmt.Printf("slice len: %d\n",len(sl))
	fmt.Printf("slice cap: %d\n",cap(sl))
	// slice 3 4 5
//	slice := s1

	fmt.Printf("----------map------------\n")
	m:=map[string]string{"key":"val"}
	m["key1"] = "val1"

	for i := 1 ; i <= 10 ; i++ {
		key_name := "key" + strconv.Itoa(i)
		val_name := "val" + strconv.Itoa(i)
		m[key_name] = val_name
	}
	fmt.Printf("map : %s\n",m)

//	fmt.Printf("----------map------------\n")
}

func test_fun() int {
	fmt.Printf("i am test func \n")
	return 0
}


