package main

import (
    "fmt"
    "net"
	"time"
)



func send_udp(host string, port int,msg string){
    // 创建连接
    socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
        IP:   net.IPv4(127, 0, 0, 1),
        Port: port,
    })
    if err != nil {
        fmt.Println("连接失败!", err)
        return
    }
    defer socket.Close()

    // 发送数据
    senddata := []byte(msg)
	_ , err = socket.Write(senddata)
    if err != nil {
        fmt.Println("发送数据失败!", err)
        return
    }
}


func main() {
	n := 1
	for {
		time.Sleep( 1 * time.Second )
		send_udp("127, 0, 0, 1",80,"786876789")
		println("conut",n)
		n = n + 1
	}
    // 接收数据
//    data := make([]byte, 4096)
//    read, remoteAddr, err := socket.ReadFromUDP(data)
//    if err != nil {
//        fmt.Println("读取数据失败!", err)
//        return
//    }
//    fmt.Println(read, remoteAddr)
//    fmt.Printf("%s\n", data)
}

