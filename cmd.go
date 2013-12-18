package main
import (
	"fmt"
	"os/exec"
	"time"
	"errors"
//	"os"
//	"log"
//	"strconv"
//	"encoding/json"
)


func main(){
	println("============================")
	r,err := ExecCmd("asad",2)
	fmt.Println(r,"err",err)
	println("============================")
	r,err = ExecCmd("pwd",2)
	fmt.Println(r,"err",err)
	println("============================")
	r,err = ExecCmd("sleep 100",1)
	fmt.Println(r,"err",err)
	println("============================")
	time.Sleep(time.Second * time.Duration(500))
}

func ExecCmd(cmd string,t int64) (r []string, err error){
	quit_chan_bool := make(chan bool)
	time_out_chan_bool := make(chan bool)
	go func(c chan bool,time_out int64){
		time.Sleep(time.Second * time.Duration(t))
		c <- true
	}(time_out_chan_bool,t)
	go func(cmd string){
		out, e1 := exec.Command("/bin/bash","-c",cmd).CombinedOutput()
		if e1 != nil{
			err = errors.New(string(out)+e1.Error())
		}else{
			r = append(r, string(out))
		}
		quit_chan_bool <- true
	}(cmd)

	select{
	case <- quit_chan_bool :
		break
	case <- time_out_chan_bool :
		err = errors.New("time out")
		return nil,err
	}

	if r == nil { return nil,err }
	return r,nil
}


