package main

//import "os"
import "fmt"
func main() {
	intarr := IntArr{2, 3, 1, -9, 0}  
	arr := int{1,4,2,8,6,3}
	a5 := int{2,7,6,3,5,8,0}
	fmt.Printf("%v\n",a5)
	fmt.Printf("%v\n",arr)
	fmt.Printf("%v\n",intarr)
}



func BubbleSort(array []int) {
	fmt.Printf("%v\n",array)
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				// 交换  
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

