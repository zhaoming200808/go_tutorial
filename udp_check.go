package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
)

type Domain struct {
	domain int
	ports_status map[int]bool
}

func NewDomain(d int,ps map[int]bool) *Domain{
	return &Domain{
		domain : d,
		ports_status  : ps,
	}
}

func get_udp_msg(port int) {
	listen_socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	})
	if err != nil {
		fmt.Println("监听失败!", err)
		return
	}
	defer listen_socket.Close()
	listen_data := make([]byte, 4096)

	Here:
//	if checkPort(port) == true {
//		return
//	}
	listen_read, listen_remoteAddr, err := listen_socket.ReadFromUDP(listen_data)
	if err != nil {
		fmt.Println("读取数据失败!", err)
		goto Here
	}
	_ = listen_remoteAddr
//	fmt.Println("\n",listen_read, listen_remoteAddr)
//	fmt.Printf("%s\n\n", listen_data)
//	fmt.Printf("%d\n\n", port)
	receive_str := string(listen_data[:listen_read])
	domian_int, err := strconv.Atoi(receive_str)
	_ = domian_int
	if err != nil {
		goto Here
	}

	setTrue(domian_int,port)
	goto Here
}


func send_udp_msg(domain string,port int,domain_str string){
	time.Sleep(100 * time.Millisecond)
	//domain_and_port localhost:8000
	domain_and_port := domain + ":" + strconv.Itoa(port)
	//println("domain_and_port",domain_and_port)
	udp_addr,err := net.ResolveUDPAddr("udp4",domain_and_port)
	//udp_addr,err := net.ResolveUDPAddr("udp4","127.0.0.1")
	//		fmt.Printf("udp_addr: |%v|",udp_addr)
	socket, err := net.DialUDP("udp4", nil, udp_addr)
	//		socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
	//			IP:   net.IPv4(127, 0, 0, 1),
	//			Port: port,
	//		})
	if err != nil {
		fmt.Println("连接失败!", err)
		return
	}
	defer socket.Close()

	// 发送数据
	senddata := []byte(domain_str)
//	domain_int := strconv.Atoi()
	for {
		_ , err = socket.Write(senddata)
		if err != nil {
			fmt.Println("发送数据失败!", err)
			return
		}
		time.Sleep(200 * time.Millisecond)
	}
}





func setTrue(d,p int){
	for _, domain := range all_domaim{
		if domain.domain == d {
			domain.ports_status[p] = true
			return
		}
	}
}

func checkDoamin(d int) bool{
	find_flag := false
	for _, domain := range all_domaim{
		if domain.domain == d {
			find_flag = true
			for _, domain_status := range domain.ports_status {
				if domain_status != true {
					return false
				}
			}
		}
	}
	if find_flag == true {
		return true
	}
	return false
}

func checkPort(p int) bool{
	for _, domain := range all_domaim {
		if domain.ports_status[p] != true {
			return false
		}
	}
	return true
}

func checkOk() bool{
	for _,domain := range all_domaim {
		//domian_num := domain.domain
		port_num_status := domain.ports_status
		for _ , port_status := range port_num_status {
			if port_status != true {
				return false
			}
		}
	}
	return true
}

func showDoaminMsg(domian_port_status map[int]bool){
	for port_num,port_status := range domian_port_status{
		if port_status != true{
			println(port_num,"NO BIAND ERROR")
		}else{
			println(port_num,"BAIND success")
		}
	}
}

func showAllDomain(){
	for _,domain := range all_domaim {
		domian_num := domain.domain
		port_num_status := domain.ports_status
		if checkDoamin(domian_num) == true{
			println(domian_num,"OK")
		}else{
			showDoaminMsg(port_num_status)
		}
	}
}


var (
	ports []int
	domains []int
	ports_status map[int]bool
	domians_status map[int]bool
	all_domaim []Domain
)

func init(){
	if ports_status == nil {
		ports_status = make(map[int]bool)
	}
	if domians_status == nil {
		domians_status = make(map[int]bool)
	}

//	ports = []int{80,8000,8001,8002,8003,8004}
//	domains = []int{4,5,6}
//	for _,port := range ports{
//		ports_status[port] = false
//	}
//	for _ , domain_num := range domains{
//		domain := *NewDomain(domain_num,ports_status)
//		all_domaim = append(all_domaim,domain)
//	}
}

func check_domain(domainNums,portNums []int,t time.Duration) {
	//get domian and port msg
	ports = portNums
	domains = domainNums
	//create domain_struct and ports_map domian_map struct
	for _,port := range ports{
		ports_status[port] = false
	}
	for _,domain := range domains{
		domians_status[domain] = false
	}

	for _ , domain_num := range domains{
		domain := *NewDomain(domain_num,ports_status)
		all_domaim = append(all_domaim,domain)
	}

	//run udp_server
	for _, port_num := range ports{
		go get_udp_msg(port_num)
	}

	//run udp_client
	//create domain_and_port 
	//localhost:80
	url_base := "localhost"
	for _,domain := range all_domaim {
		domian_num := domain.domain
		port_num_status := domain.ports_status
		domian_str := strconv.Itoa(domian_num)
		url := url_base + domian_str
		for port_num , _ := range port_num_status {
			go send_udp_msg(url,port_num,domian_str)
		}
	}

    tick := time.Tick(50 * time.Millisecond)
    boom := time.After(t * time.Millisecond)
    for{
        select {
        case <- tick :
			if checkOk() == true {
				println("all domian baind success")
				return
			}
            fmt.Printf("tick :%d\n",50)
        case <- boom :
            fmt.Printf("boom :%d\n",3000)
			showAllDomain()
            return
        }
    }

}

func main() {
	p := []int{80,8000,8001,8002,8003,8004}
	d := []int{4,5,6}
	println("===========")
	check_domain(d,p,3000)
	println("===========")
//	fmt.Printf("%v\n",all_domaim)
//	url_base := "localhost"
//	port_nums := []int{80,8000,8001,8002,8003,8004}
//
//	for _, port_num := range port_nums{
//		go get_udp_msg(port_num)
//	}
//	for _,domain := range all_domaim {
//		domian_num := domain.domain
//		port_num_status := domain.ports_status
//		domian_str := strconv.Itoa(domian_num)
//		url := url_base + domian_str
//		for port_num , _ := range port_num_status {
//			go send_udp_msg(url,port_num,domian_str)
//		}
//	}
//	time.Sleep(1 * time.Second)
//	if checkOk() == true {
//		println("OKOKOK!!!")
//	}else{
//		println("ONNNNNNNNNNNN")
//	}
//	showAllDomain()
//	fmt.Printf("%v\n",all_domaim)
}


