package main

import (
	"database/sql"
	"fmt"
	memcache "github.com/bradfitz/gomemcache/memcache"
	_ "github.com/go-sql-driver/mysql"
	redis "github.com/gosexy/redis"
	fdfs "github.com/weilaihui/fdfs_client"
	mgo "gopkg.in/mgo.v2"
	"os"
	"runtime"
	"strconv"
	"time"
)

func CheckErr(err error, operating string) {
	//pc,file,line,ok = runtime.Caller(1)
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	if err != nil {
		fmt.Printf("@@@ ERROR: |%s| %s failed.\n", funcName, operating)
		fmt.Printf("  %s\n", err.Error())
		//panic(err)
		os.Exit(-1)
	}
	fmt.Printf("### OK: |%s| %s success.\n", funcName, operating)
}

func mysql_check(host string, port int, username string, password string, database string) {
	conn_str := username + ":" + password + "@" + "tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database + "?charset=utf8"
	//fmt.Printf("%s\n", conn_str)
	//db, err := sql.Open("mysql", "replication_user:spritereplication@tcp(192.168.133.187:3306)/test?charset=utf8")
	db, err := sql.Open("mysql", conn_str)
	defer db.Close()
	CheckErr(err, "mysql ping host")
	rows, err := db.Query("select 1 as num")
	CheckErr(err, "mysql login and select")
	var num int
	for rows.Next() {
		err = rows.Scan(&num)
		CheckErr(err, "mysql get value")
	}
	if num != 1 {
		fmt.Printf("ERROR: mysql num != 1")
	}
}

func fdfs_check(host string, port int) {
	var hostList = []string{host}
	filePath := "./client.conf"
	tracker := &fdfs.Tracker{
		HostList: hostList,
		Port:     port,
	}
	fdfsClient, err := fdfs.NewFdfsClientByTracker(tracker)
	CheckErr(err, "create fdfsClient")

	uploadResponse, err := fdfsClient.UploadByFilename(filePath)
	CheckErr(err, "upload file "+filePath)

	//	fmt.Println(uploadResponse.GroupName)
	//  fmt.Println(uploadResponse.RemoteFileId)

	fdfsClient.DeleteFile(uploadResponse.RemoteFileId)
	CheckErr(err, "delete file "+filePath)
}

func mgo_check(host string, port int) {
	//time_str := time.Now().Format("2006-01-02 15:04:05")
	session, err := mgo.Dial(host + ":" + strconv.Itoa(port))
	defer session.Close()
	CheckErr(err, "Dial "+host+":"+strconv.Itoa(port))
	// Optional. Switch the session to a monotonic behavior.
	//session.SetMode(mgo.Monotonic, true)

	c := session.DB("admin").C("user")
	//err = c.Insert(&Check{time_str})
	num, err := c.Count()
	CheckErr(err, "selec count(admin.user:"+strconv.Itoa(num)+")")
}

func mc_check(host string, port int) {
	time_str := time.Now().Format("2006-01-02 15:04:05")
	key_str := "op_check_time"
	mc := memcache.New(host + ":" + strconv.Itoa(port))
	err := mc.Set(&memcache.Item{Key: key_str, Value: []byte(time_str)})
	CheckErr(err, "set value("+key_str+":"+time_str+")")
	it, err := mc.Get(key_str)
	CheckErr(err, "get value("+key_str+":"+time_str+")")
	if time_str != string(it.Value) {
		fmt.Printf("ERROR: |main.mc_check| get value error\n")
	}

}

func redis_check(host string, port uint, auth string) {
	//	var client *redis.Client
	time_str := time.Now().Format("2006-01-02 15:04:05")
	key_str := "op_check_time"
	client := redis.New()
	err := client.Connect(host, port)
	CheckErr(err, "conn redis")
	defer client.Quit()
	client.Auth(auth)
	s, err := client.Ping()
	CheckErr(err, "send ping "+s)
	ret, err := client.Set(key_str, time_str)
	CheckErr(err, "set value("+key_str+":"+time_str+")")
	if ret != "OK" {
		fmt.Printf("ERROR: |main.redis_check| ret error\n")
	}
	value, err := client.Get(key_str)
	CheckErr(err, "get value("+key_str+":"+time_str+")")
	if value != time_str {
		fmt.Printf("ERROR: |main.redis_check| get value error\n")
	}
}

func main() {
	fmt.Printf("==================================\n")
	mysql_check("192.168.133.187", 3306, "replication_user", "spritereplication", "test")
	fmt.Printf("==================================\n")
	redis_check("192.168.133.183", 6079, "85fsv#wgJmLgJQhyQfsyOJ11l1xiG3XZ")
	fmt.Printf("==================================\n")
	mc_check("192.168.133.84", 11211)
	fmt.Printf("==================================\n")
	mgo_check("192.168.133.72", 30001)
	fmt.Printf("==================================\n")
	fdfs_check("duke", 33133)
	fmt.Printf("==================================\n")
	n := add(1, 2)
	print(n)
}
