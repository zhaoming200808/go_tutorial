package main

import (
    "fmt"
	"reflect"
	"strconv"
)

func prints(i int) string {
	fmt.Println("i =",i)
	return strconv.Itoa(i)
}

func main() {
	fv := reflect.ValueOf(prints)
	params := make([]reflect.Value,1)  //参数
	params[0] = reflect.ValueOf(50)   //参数设置为20
	rs := fv.Call(params)              //rs作为结果接受函数的返回值
	fmt.Println("result:",rs[0].Interface().(string)) //当然也可以直接是rs[0].Interface()
}





