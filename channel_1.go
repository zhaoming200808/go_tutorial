package main
import "fmt"
import "time"
func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}
func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

//	for n , ch := range(chs) {
//		println(n)
//		<-ch
//	}

	// 首先，我们实现并执行一个匿名的超时等待函数
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // 等待1秒钟
		timeout <- true
	}()
	// 然后我们把timeout这个channel利用起来
	n := 0
	select {
	case <- chs:
		n = n + 1
		println(n)
	case <-timeout:
		return
		// 一直没有从ch中读取到数据，但从timeout中读取到了数据
	}

}
