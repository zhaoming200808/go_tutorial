package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func CheckErr(err error, operating string) {
	//pc,file,line,ok = runtime.Caller(1)
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	if err != nil {
		log.Printf("@@@ ERROR: |%s| %s failed.\n", funcName, operating)
		log.Printf("  %s\n", err.Error())
		//panic(err)
		os.Exit(-1)
	}
	log.Printf("### OK: |%s| %s success.\n", funcName, operating)
}

type T struct {
	createTime time.Time
	startTime  time.Time
	endTime    time.Time
}

func NewT() *T {
	return &T{
		createTime: time.Now(),
	}
}
func (this *T) Show() {
	fmt.Printf("createTime:%v\n", this.createTime)
}

func init() {

}

func Work(num int, c chan int) (n int) {
	log.Printf("work : %d\n", num)
	time.Sleep(1 * time.Second)
	<-c
	return num
}

func Handler() {
	maxWork := 5
	c := make(chan int, maxWork)

	works := 15
	for i := 1; i <= works; i++ {
		c <- 1
		go Work(i, c)
	}
	for {
		if len(c) == 0 {
			println(len(c))
			log.Printf("Handler down")
			break
		} else {
			println(len(c))
			time.Sleep(1 * time.Second)
			//time.Sleep(100 * time.Millisecond)
			log.Printf("sleep 1 sec")
		}
	}
}

func main() {
	fmt.Printf("==============================================================\n")
	Handler()
	log.Printf("down")

	fmt.Printf("==============================================================\n")
}
