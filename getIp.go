package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("udp", "google.com:80")
	defer conn.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])

}

