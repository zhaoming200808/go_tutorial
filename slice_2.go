package main

import "fmt"

func copy_map(oldMap map[int]bool) (newMap map[int]bool){
	newMap = make(map[int]bool)
	for k,v := range oldMap{
		newMap[k] = v
	}
	return
}

func main() {
	m := map[int]bool{80:false,8000:false,8001:false,8002:false,8003:false,8004:false}
//	po := []int{80,8000,8001,8002,8003,8004}
	i := []int{4,5,6}
	var sl []map[int]bool
	fmt.Printf("%v\n",sl)

	for _,_ = range i{
//		tmp_m := copy_map(m)
		sl = append(sl,copy_map(m))
	}
	fmt.Printf("%v\n",sl)
//	fmt.Printf("%v\n",m)


	fmt.Printf("%v\n",sl)

	sl[0][80]=true

	fmt.Printf("%v\n",sl)

	return


    p := []int{2, 3, 5, 7, 11, 13}

    for i := 0; i < len(p); i++ {
        fmt.Printf("p[%d] == %d\n",
            i, p[i])
    }

    fmt.Println("--------------------------")
    fmt.Println("p ==", p)
    fmt.Println("p[1:4] ==", p[1:4])

    // 省略下标代表从 0 开始
    fmt.Println("p[:3] ==", p[:3])

    // 省略上标代表到 len(s) 结束
    fmt.Println("p[4:] ==", p[4:])

	slice_msg := `
	s[lo:hi]		//表示从 lo 到 hi-1 的 slice 元素
	s[lo:lo]		//是空的
	s[lo:lo+1]		//有一个元素
	`
	fmt.Println(slice_msg)

    a := make([]int, 5)
    printSlice("a", a)
    b := make([]int, 0, 5)
    printSlice("b", b)
    c := b[:2]
    printSlice("c", c)
    d := c[2:5]
    printSlice("d", d)

	slice_make_msg := `
	为了指定容量，可传递第三个参数到 'make'：
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
	`
	fmt.Println(slice_make_msg)

	var z []int
    fmt.Println(z, len(z), cap(z))
    if z == nil {
        fmt.Println("nil!")
    }

}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}




