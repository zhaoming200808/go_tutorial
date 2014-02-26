package main

import "fmt"
import "os"
import "bytes"
import "os/exec"
import "time"
import "errors"
func main() {
//	bey , err:= run("/tmp","/bin/bash","-c","sleep 100")
//	var time_out time.Duration
//	time_out = 1000
	bey , err:= run(".",100,"ls -al")
	if err != nil {
		fmt.Printf("return:|%s|",string(bey))
		println("-----------------")
		fmt.Printf("err:%s\n",err.Error())
	}
	fmt.Printf("return:|%s|",string(bey))
	println("\r\n-----------------\r\n")
	fmt.Println("end")

//	time.Sleep(10 * time.Second)
}


func run(dir string,t time.Duration ,c ...string) (v []byte, err error) {
	//check dir
	file, err := os.Stat(dir)
	if err != nil || file.IsDir() != true{
		fmt.Println("%s",err.Error())
		return nil,errors.New("path not find: " + dir)
	}
	//check args
	args := make([]string,0,10)
	args = append(args,"/bin/bash","-c")
	for _,v := range c{
		args = append(args,v)
	}

	quit := make(chan bool)
	value := make(chan string)
	var buf bytes.Buffer
	var cmd *exec.Cmd
	go func(quit chan bool){
		time.Sleep(t * time.Millisecond)
		quit <- true
	}(quit)

	go func(value chan string){
		cmd = exec.Command(args[0], args[1:]...)
		cmd.Dir = dir
		cmd.Stdout = &buf
		cmd.Stderr = cmd.Stdout
		err = cmd.Run()
		cmd.ProcessState.Exited()
		v = buf.Bytes()
		value <- "ok"
	}(value)
	select {
	case <-value:
		return v, err
	case <-quit:
		if cmd != nil {
			fmt.Printf("%v\n",cmd.Process.Kill())
			return nil,errors.New("run cmd time out")
		}
		return nil,nil
	}
}


