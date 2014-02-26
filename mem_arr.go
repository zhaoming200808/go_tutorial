package main

//import "os"
import "fmt"
import "strconv"
import "time"
import "thrift/rpc"

func test_arr_mem() {
	fmt.Printf("---------array------------\n")
	var arr [1000000]string
	println(len(arr))
	for i := 0 ; i < len(arr) ; i++{
		value := "value" + strconv.Itoa(i)
		arr[i] = value
	}
	println(len(arr))
	time.Sleep(60000 * time.Millisecond)
	if true { return }
}


func test_slice_mem() {
	println("stret:",time.Now().Unix())
	fmt.Printf("---------array------------\n")
	sl := make([]string, 1000000)
//	for i := 0 ; i < 1000000 ; i++{
//		value := "value" + strconv.Itoa(i)
//		sl=append(sl,value)
//	}
	println("init ok:",time.Now().Unix())
	for i , value := range sl{
		if i != 1000000 {
			continue
		}
		if value == "nil" {
			println(value)
		}
	}
	println("end:",time.Now().Unix())
	time.Sleep(60000 * time.Millisecond)
	if true { return }
}

type Host struct {
    cmd_chan chan string
    flag_str string
    host_addr string
    conn_flag bool
	client *rpc.RpcServiceClient
    err error
}

func NewHost(flag,host_addr string) *Host {
    if flag == "" || host_addr == "" {
        return nil
    }
    return &Host{
        cmd_chan:   make(chan string),
        flag_str:   flag,
        host_addr:  host_addr,
    }
}

func main() {

	//test_arr_mem()
	//test_slice_mem()

//	var arr [100000]*Host
//	println(len(arr))
//	for i := 0 ; i < len(arr) ; i++{
//		flag := "flag" + strconv.Itoa(i)
//		host := NewHost(flag,"255.255.255.255:65535")
//		arr[i] = host
//	}


	sl := make([]*Host, 1000000)
	for i := 0 ; i < 1000000 ; i++{
		flag := "flag" + strconv.Itoa(i)
		host := NewHost(flag,"255.255.255.255:65535")
		sl=append(sl,host)
	}

//	fmt.Printf("%v\n",arr)
	println("ok")
	time.Sleep(60000 * time.Millisecond)
	if true { return }


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



