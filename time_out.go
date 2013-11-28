package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {  //进入 无限循环状态 
        select { //当下面有case条件满足时就运行与之对应的语句 当多个满足时 随机选择一个
		//当c管道有数据时 就生成一个符合规范的数 并返回
        case c <- x:
            x, y = y, x+y
		//当quit管道有数据时 就打印并退出
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}


type Sync struct {
	go_count int

}


func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 30; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}



