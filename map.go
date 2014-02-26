package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }

//	var m1 = map[string]string{"a":"b"}
//	m1:=map[string]string{"key":"val"}
	m1 := make(map[string]string)
	m1["jian"] = "value"
    fmt.Println(m1["jian"])

    fmt.Println(m["Bell Labs"])

	map_init_msg := `
	map声明 定义方式: （map使用前必须make）
	1:
	  var m map[string]string					//声明
	  m = make(map[string]string)				//定义(make )
	2:
	  m = make(map[string]string)				//声明并定义
	3:
	  var m = map[string]string{"key":"val"}	// 相当于隐式make

	`
	fmt.Println(map_init_msg)

	map_operating_msg := `
	map 操作：
	m[key] = elem		//插入或修改一个元素
	elem = m[key]		//获得元素
	delete(m, key)		//删除元素

	elem, ok = m[key]	//通过双赋值检测某个键存在
	如果 key 在 m 中，'ok' 为 true 。否则， ok 为 'fals'，并且 elem 是 map 的元素类型的零值。
	同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。
	`

	fmt.Println(map_operating_msg)


	mm := make(map[int]map[int]bool)
	println(mm)
























}




