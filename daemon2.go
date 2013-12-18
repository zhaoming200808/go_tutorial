package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main(){

//	if os.Getppid()!=1{           //判断当其是否是子进程，当父进程return之后，子进程会被 系统1 号进程接管
//		filePath,_:=filepath.Abs(os.Args[0])  //将命令行参数中执行文件路径转换成可用路径
//		cmd:=exec.Command(filePath,os.Args[1:]...)//将其他命令传入生成出的进程
//		cmd.Stdin=os.Stdin                               //给新进程设置文件描述符，可以重定向到文件中
//		cmd.Stdout=os.Stdout
//		cmd.Stderr=os.Stderr
//		cmd.Start()                                          //开始执行新进程，不等待新进程退出
//		return
//	}
	if os.Getppid() !=1 {
		filePath,_:=filepath.Abs(os.Args[0])  //将命令行参数中执行文件路径转换成可用路径
		args:=append([]string{filePath},os.Args[1:]...)
		os.StartProcess(filePath,args,&os.ProcAttr{Files:[]*os.File{os.Stdin,os.Stdout,os.Stderr}})
		return
	}


//	for i:=1 ; i<= 30 ; i++{
//		println(i)
//		time.Sleep(1 * time.Second)
//	}

	for {
		print("true")
		time.Sleep(1 * time.Second)
	}


}


func daemon(){
	if os.Getppid()!=1{           //判断当其是否是子进程，当父进程return之后，子进程会被 系统1 号进程接管
		filePath,_:=filepath.Abs("daemon2")  //将命令行参数中执行文件路径转换成可用路径
		cmd:=exec.Command(filePath,os.Args[1:]...)//将其他命令传入生成出的进程
		cmd.Stdin=os.Stdin                               //给新进程设置文件描述符，可以重定向到文件中
		cmd.Stdout=os.Stdout
		cmd.Stderr=os.Stderr
		cmd.Start()                                          //开始执行新进程，不等待新进程退出
		return
	}
}

