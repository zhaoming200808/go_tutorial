package main

import "fmt"
import "strconv"
import "time"
import "runtime"

type IHost interface {
	Exec(cmd string) Value
}

type Host struct {
	host_cmd_chan chan string
	host_flag_str string
//	host_ip	string
//	game_id int64
//	group_id int64
//	online_flag bool
}

func NewHost(flag string) *Host {
	return &Host{
		host_cmd_chan:	make(chan string),
		host_flag_str:	flag,
	}
}

type Value struct {
	host_flag_str string
	cmd_return_err string
}

func (h Host) Exec(cmd string) (v Value){
	time.Sleep(10 * time.Millisecond)
//	fmt.Printf("%s exec: %s\r\n",h.host_flag_str,cmd);
	v.host_flag_str = "ok"
	v.cmd_return_err = "ok"
	return
}

func Exec_cmd(i IHost,cmd string,c chan Value) (v Value){
	v = i.Exec(cmd)
	defer func(){
		c <- v
	}()
	return
}

func init() {
	println("CPU num: ",runtime.NumCPU())
	println("GOMAXPROCS num: ",runtime.NumCPU()*50)
	runtime.GOMAXPROCS(runtime.NumCPU()*50)

}

func main() {

	all_host := []Host{}
	var host Host
	var flag string
	//init
	for i := 1 ; i <= 10 ; i++ {
		flag = "f"+ strconv.Itoa(i)
		host = Host{make(chan string),flag}
		//host = NewHost(flag)
		all_host = append(all_host,host)
	}
	fmt.Printf("%v\n",all_host)

	return_value := make(chan Value,1000)

	//exec 
	for index, value := range all_host {
		if value.host_flag_str != "f4" {
			fmt.Printf("%d  %v EXEC\n",index,value)
			go Exec_cmd(value,"pwd",return_value)
		}else {
			fmt.Printf("%d  %v NO EXEC!\n",index,value)
			time.Sleep(10 * time.Millisecond)
		}
	}

	//
//	time.Sleep(10 * time.Millisecond)

	r := Value{}
	tick := time.Tick(200 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)
	for{
		select {
		case  return_value <- r:
			if r.host_flag_str != "" {
				fmt.Printf("v :%v\n",r)
			}else{
				return
			}
		case <- tick :
			fmt.Printf("tick :%d\n",200)
		case <- boom :
			fmt.Printf("boom :%d\n",1000)
			return
		}
	}
}





