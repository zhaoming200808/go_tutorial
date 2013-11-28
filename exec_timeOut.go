package main

import (
	"fmt"
//	"log"
	"os/exec"
	"time"
//	"strings"
)




func ecex_cmd(cmd_string string,t time.Duration) (r []string, err error){
	ok := make(chan bool)
	time_out := make(chan bool)
	//set time_out
	go func(c chan bool,time_out time.Duration){
		time.Sleep(time.Second * time.Duration(time_out))
		c <- true
	}(time_out,t)

	//exec cmd
	go func(cmd_string string){
		out, t_err := exec.Command("/bin/bash","-c",cmd_string).Output()
		if t_err != nil {
			fmt.Printf("%s\n",t_err.Error())
			//log.Fatal(err)
			err = t_err
		}
//		fmt.Printf("The date is %s\n", out)
		r = append(r, string(out))
		ok <- true
	}(cmd_string)

	//wait cmd done or time-out
	select{
	case  <- ok:
		println("ok")
	case  <- time_out:
		println("time_out")
	}
//	fmt.Printf("The date is %s,%v\n", out,c)
	return
}


func main() {
	a,b := ecex_cmd("sleep 10",5)
	fmt.Printf("%v,%v\n",a,b)
}

