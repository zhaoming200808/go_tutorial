package main

//import "os"
import "fmt"
import "strconv"
func main() {
	var s1 string // = "a"
	//var s1 = "";
	s1 = "i am s1"
	s2 := "i am s2"
	s3 := "i am s3"

	fmt.Printf("%s\n",s1)
	fmt.Printf("%s\n",s2)
	fmt.Printf("%s\n",s3)

	var i1 int // = 1
	i1 = 1
	i2 := 2
	i3 := 3
	fmt.Printf("%d\n",i1)
	fmt.Printf("%d\n",i2)
	fmt.Printf("%d\n",i3)

	//int 整形
	s_i := s1 + strconv.Itoa(i1)
	//float64 浮点型
	strconv.FormatFloat(float64(2.7),'f', -1, 32)
	//int16 uint int32 类整形
	strconv.Itoa(int(1))

	fmt.Printf("%s\n",s_i)

	//str 
	str1 := "i am str1 \n" //可转义字符串
	str2 := `i am str2 \n` //绝对字符串
	fmt.Printf("%s\n",str1)
	fmt.Printf("%s\n",str2)
	//print
	//i am str1 
	//i am str2 \n
	//
}




