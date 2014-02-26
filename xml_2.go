package main

import (
	"encoding/xml"
	"fmt"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"io/ioutil"
	"strconv"
	"net/http"

//	"io"
	"time"
	"os"
	"math/rand"
)

type Hoolai struct {
	XMLName xml.Name `xml:"Hoolai"`
	SuperServer	string	`xml:"SuperServer"`
	LoginServer	string	`xml:"LoginServer"`
	SessionServer	string	`xml:"SessionServer"`
	ScenesServer	string	`xml:"ScenesServer"`
	GatewayServer	string	`xml:"GatewayServer"`
	Tgw		string	`xml:"global>tgw"`
	ServerGroupId	int	`xml:"global>serverGroupId"`
	Billserverip	string	`xml:"global>billserverip"`
	Biserverip	string	`xml:"global>biserverip"`
	Goodsurl	string	`xml:"global>goodsurl"`
	UseTencentWordFilter	string	`xml:"global>useTencentWordFilter"`
	UseLocalWordFilter	string	`xml:"global>useLocalWordFilter"`
	Log	string	`xml:"global>log"`
	ThreadPoolCapacity	string	`xml:"global>threadPoolCapacity"`
	Mysql	string	`xml:"global>mysql"`
	Ifname	string	`xml:"global>ifname"`
	Superserver	string	`xml:"global>superserver"`
}

//golbal var
var (
	v Hoolai //config_xml_status
	xml_config_filePath string
	xml_loginzonelist_filePath string
	ports []int //server_ports
	domian_nums map[int]bool
	zone_nums map[int]bool
	mysql_conn_str string
)

func init() {
	xml_config_filePath = "/usr/local/services/server/bin/config.xml"
	xml_loginzonelist_filePath = "/usr/local/services/server/bin/loginzonelist.xml"
	ports = append(ports,8000,8001,8002,8003,8004)
	domian_nums = make(map[int]bool)
	zone_nums = make(map[int]bool)
}
func main() {
	fmt.Printf("==============CHECK==================\n")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Runtime error caught: %v\n", r)
		}
	}()
	//check xml file
	println("-------------check xml file----------------------")
	mysql_conn_str = test_config_xml(xml_config_filePath)
	test_loginzonelist_xml(xml_loginzonelist_filePath)
	println("-----------------------------------------------")
	fmt.Printf("%v\n",domian_nums)
	println("-------------check mysql conn----------------------")
	//check mysql
	test_mysql(mysql_conn_str)
	println("----------!--check domain conn--!-------------------")
	test_ports_and_domain()
	log.Println(v.ServerGroupId)
	println("-------------show server msg------------------------")
	show_server_msg()
	fmt.Printf("==============CHECK OK================\n")
}

func show_server_msg(){
	var domian_nums_str string
	for k,_ :=  range domian_nums {
		str_domian_num := strconv.Itoa(k)
		if domian_nums_str == "" {
			domian_nums_str = str_domian_num
		}else {
			domian_nums_str = domian_nums_str + " " + str_domian_num
		}
	}
	fmt.Printf("======\tgroupid: |%d|\n",v.ServerGroupId)
	fmt.Printf("======\tdomains: |%s|\n",domian_nums_str)
	fmt.Printf("======\tmysql:   |%s|\n",v.Mysql)
}

func test_ports_and_domain(){
	for _ , port :=  range ports {
		go listen_port(port)
	}
	//sleep

	//create magic num
	seed := time.Now().UnixNano() + int64(v.ServerGroupId) + int64(os.Getpid()) + int64(os.Getuid()) + int64(os.Getppid())
	r := rand.New(rand.NewSource(seed))
	magic_num := r.Int()
	println("magic_num",magic_num)
	//check
	time.Sleep(1 * time.Second)
	for k , _ :=  range domian_nums {
		domain := "s" + strconv.Itoa(k) + ".app100715380.qqopenapp.com"
		domain = "localhost"
		for _ , port :=  range ports {
			println(k,port,domain)
//			l , err := net.Dial("tcp", domain + ":" + strconv.Itoa(port))
//			if err == nil {
//				println("conn",domain + ":" + strconv.Itoa(port) , "OK")
//			}else {
//				log.Fatal("error",err)
//			}
//			defer l.Close()
//			l.Write([]byte(strconv.Itoa(k)))
		}
	}

	time.Sleep(4 * time.Second)
}

func listen_port(k int){
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Runtime error caught: %v\n", r)
		}
	}()
//	var map[string]int
	println("listen:",strconv.Itoa(k))
	http.HandleFunc("/", check)
	http.ListenAndServe(":" + strconv.Itoa(k), nil)
}

func check(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

type Zone struct {
	Id int `xml:"id,attr"`
}

type LoginZoneList struct {
	XMLName xml.Name `xml:"Hoolai"`
	Test string `xml:"test"`
	Zone []Zone `xml:"zone"`
}

func test_loginzonelist_xml(xml_filePath string){
	//log.Println(xml_filePath)
	v2 := LoginZoneList{}
	xmlFile, err := ioutil.ReadFile(xml_filePath)
	data := string(xmlFile)
	//fmt.Printf("%s",data)
	err = xml.Unmarshal([]byte(data), &v2)
	if err != nil {
		//fmt.Printf("error: %v", err)
		log.Fatal(err)
	}
	for _ , zone := range v2.Zone {
		//fmt.Printf("\nzone_num is|%d|\n",zone.Id)
		domian_nums[zone.Id] = false
	}
	log.Printf("TEST XML loginzonelist SUCCESS")
}

func test_config_xml(xml_filePath string) (mysql_conn_str string){
	v = Hoolai{}
	xmlFile, err := ioutil.ReadFile(xml_filePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		log.Fatal(err)
	}
	//fmt.Printf("%v",string(xmlFile))
	data := string(xmlFile)
	err = xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		log.Fatal(err)
	}
	mysql_con_str := v.Mysql
	//Replace 
	mysql_con_str = strings.Replace(mysql_con_str,"mysql://","",1)
	mysql_con_str = strings.Replace(mysql_con_str,"@"," ",1)
	mysql_con_str = strings.Replace(mysql_con_str,"/"," ",1)
	mysql_con_str = strings.Replace(mysql_con_str,":"," ",2)
	//split args
	mysql_con_str_arr := strings.SplitN(mysql_con_str," ",-1)
	//check
	if len(mysql_con_str_arr) != 5 {
		println("error")
		log.Fatal("mysql_con_str_arr error")
	}
	mysql_user := mysql_con_str_arr[0]
	mysql_passwd := mysql_con_str_arr[1]
	mysql_host := mysql_con_str_arr[2]
	mysql_port := mysql_con_str_arr[3]
	mysql_dbName := mysql_con_str_arr[4]
	mysql_conn_str = mysql_user + ":" + mysql_passwd + "@tcp(" + mysql_host + ":" + mysql_port + ")/" + mysql_dbName
	domian_nums[v.ServerGroupId] = true
	log.Printf("TEST XML config.xml SUCCESS")
	return mysql_conn_str
}

func test_mysql(mysql_conn_str string){
	db, err := sql.Open("mysql", mysql_conn_str)
	if err != nil {
		log.Println(err)
		log.Fatal(`TEST DB NETWORK LINK FAILED`)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
		log.Fatal(`TEST DB CONN FAILED`)
	}else{
		log.Println(`TEST DB CONN SUCCESS`)
	}

	rows, err := db.Query("select 1")
	if err != nil {
		log.Println(err)
		log.Fatal(`TEST DB Query FAILED`)
	}else {
		log.Println(`TEST DB Query SUCCESS`)
	}
	defer rows.Close()
	var t_num int
	for rows.Next(){
		if err := rows.Scan(&t_num); err != nil {
			log.Println(err)
			log.Fatal(`TEST DB GET RESULT FAILED`)
		}
	}
	if t_num != 1 {
		log.Fatal(`TEST DB SELECT 1 FAILED`)
	}else {
		log.Println(`TEST DB SELECT 1 SUCCESS`)
	}
}

