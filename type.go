package main

import (
    "fmt"
    "math/cmplx"
)

var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)


func main() {
    const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)

	base_type := `
Go 指南
基本类型
Go 的基本类型有Basic types
	- bool
	- string
	- int  int8  int16  int32  int64
	- uint uint8 uint16 uint32 uint64 uintptr
	- byte // uint8 的别名
	- rune // int32 的别名
// 代表一个Unicode码
	- float32 float64
	- complex64 complex128
`

    fmt.Printf(base_type)



	var v interface{}
//	v = [10]int{1,2,3,4,5,6,7,8,9,10}
	v = [10]string{"i am password",}
//	v = 1182837279
//	v = 2.6
//	v = 3e8
//	v = a
	switch vt := v.(type){
	case byte:
		println("byte uint8")
	case int32:
		println("rune int32")
	case string:
		println("string")
	case int,int8,int16,int64:
		println("int")
	case uint,uint16,uint32,uint64,uintptr:
		println("uint")
	case float32,float64:
		println("float")
	case complex64,complex128:
		println("complex")
	case  []interface{},interface{}:
		println("interface")
	default:
		println("no find")
		println(vt)
	}
}







