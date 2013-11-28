package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
var pow_2 = []uint64{}


func main() {
	for i:=0 ; i<=50 ; i++ {
		if i == 0 {
			pow_2 = append(pow_2,1)
		}
		var x uint64 = 2 << uint(i)
		pow_2 = append(pow_2,x)
	}

    for i, v := range pow_2 {
        fmt.Printf("2**%d = %d\n", i, v)
    }

	range_msg := `
	可以将值赋值给 _ 来忽略序号和值。
	如果只需要索引值，去掉“, value”的部分即可。
	eg:
		for _, value := range pow {
			fmt.Printf("%d\n", value)
		}
	`
	fmt.Println(range_msg)
}



