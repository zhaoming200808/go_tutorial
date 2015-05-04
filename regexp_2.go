package main

import (
	"fmt"
	"log"
	"os"

	"regexp"
	"runtime"
	"strings"
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

//  1,|api.budejie.com|
//  2,|202.99.89.144|
//  3,|-|
//  4,|-|
//  5,|[12/Oct/2014:05:10:29 +0800]|
//  6,|"GET /api/api_open.php?market=xiaomi&udid=864644027267319&a=praise&c=comment&os=4.4.2&client=android&id=6799732&per=5&visiting=10606684&mac=64%3Ab4%3A73%3A3d%3Af6%3A23&ver=4.0.6&maxtime=0 HTTP/1.1"|
//  7,|200|
//  8,|666|
//  9,|"-"|
//  10,|"-"|
//  11,|"-"|
//  12,||
//  13,|200|
//  14,|0.033|
//  15,|192.168.133.186:9000|
//  16,|0.033|

func main() {
	str := `api.budejie.com 202.99.89.144 - - [12/Oct/2014:05:10:29 +0800] "GET /api/api_open.php?market=xiaomi&udid=864644027267319&a=praise&c=comment&os=4.4.2&client=android&id=6799732&per=5&visiting=10606684&mac=64%3Ab4%3A73%3A3d%3Af6%3A23&ver=4.0.6&maxtime=0 HTTP/1.1" 200 666 "-" "-" "-"  200 0.033 192.168.133.186:9000 0.033`
	str = `api.budejie.com 111.227.170.108 - - [30/Dec/2014:03:23:05 +0800] "GET /api/api_open.php?a=stat&appname=baisibudejie_NEW&asid=01EBE873-7523-4509-95A4-F11DDE32AC78&c=video&client=iphone&device=iPhone%204&flag=0&jbk=0&mac=&market=&openudid=af75047cac751f333ccf99a5875d78ccdb74fecb&pid=7813337&udid=&userid=&ver=3.1 HTTP/1.1" 200 51 "-" "MyWeiboJingXuan/3.1 (iPhone; iOS 7.1.2; Scale/2.00)" "-"  200 0.005 192.168.133.71:9000 0.005`
	str = ``
	str = `api.budejie.com 222.219.152.70 - - [31/Dec/2014:03:12:03 +0800] "GET /api/api_open.php?market=360zhushou&maxid=1419930800&udid=352246065139324&a=newlist&c=data&os=4.4.2&client=android&userID=&page=1&per=20&visiting=&type=29&time=week&mac=2c%3A8a%3A72%3A3b%3Ab1%3A87&ver=3.9.3 HTTP/1.1" 500 186 "-" "-" "-"  - - - 1.970`
	n := 100
	fmt.Printf("n: |%d|\n", n)
	fmt.Printf("str:|%s|\n", str)
	re, err := regexp.Compile(`(?U)(^.*) (.*) (.*) (.*) (\[.*\]) (\".*\") (.*) (.*) (\".*\") (\".*\") (\".*\") (.*) (.*) (.*) (.*) (.*)$`)
	CheckErr(err, "init re")
	//	fmt.Printf("re_count:%d\nstr:|%s|\n", len(re.FindStringSubmatch(str)), re.FindStringSubmatch(str)[0])
	lineArr := re.FindStringSubmatch(str)
	for i, v := range lineArr {
		fmt.Printf("%d,|%s|\n", i, v)
	}

	//  1,|api.budejie.com|
	//  2,|111.227.170.108|
	//  3,|-|
	//  4,|-|
	//  5,|[30/Dec/2014:03:23:05 +0800]|
	//  6,|"GET /api/api_open.php?a=stat&appname=baisibudejie_NEW&asid=01EBE873-7523-4509-95A4-F11DDE32AC78&c=video&client=iphone&device=iPhone%204&flag=0&jbk=0&mac=&market=&openudid=af75047cac751f333ccf99a5875d78ccdb74fecb&pid=7813337&udid=&userid=&ver=3.1 HTTP/1.1"|
	//  7,|200|
	//  8,|51|
	//  9,|"-"|
	//  10,|"MyWeiboJingXuan/3.1 (iPhone; iOS 7.1.2; Scale/2.00)"|
	//  11,|"-"|
	//  12,||
	//  13,|200|
	//  14,|0.005|
	//  15,|192.168.133.71:9000|
	//  16,|0.005|

	println(len(lineArr))
	if len(lineArr) != 17 {
		println("ERROR: len(lineArr) != 16")
		os.Exit(1)
	}
	//请求信息
	requesHost := lineArr[1]
	requesClientIp := lineArr[2]
	requesTime, err := time.Parse("[02/Jan/2006:15:04:05 +0800]", lineArr[5])
	CheckErr(err, "Parse nginx time")
	requesUserAgent := lineArr[10]

	var backendHost string
	var backendPort string
	if lineArr[15] != "-" {
		backendHostAndPort := strings.Split(lineArr[15], ":")
		if len(backendHostAndPort) != 2 {
			backendHost = ""
			backendPort = ""
		} else {
			backendHost = backendHostAndPort[0]
			backendPort = backendHostAndPort[1]
		}
	}

	requestAll := strings.Split(strings.Replace(lineArr[6], "\"", "", -1), " ")
	requestMethod := requestAll[0]
	requesUrl := requestAll[1]
	requesHttpVersion := requestAll[2]

	reAflag, err := regexp.Compile(`(?U)[\&|\?]a=(.*)[\&|$]`)
	CheckErr(err, "init reAflag")
	reCflag, err := regexp.Compile(`(?U)[\&|\?]c=(.*)[\&|$]`)
	CheckErr(err, "init reCflag")
	//	fmt.Printf("===========================\n%v\n", reAflag.FindStringSubmatch(requesUrl))
	//	fmt.Printf("===========================\n%v\n", reCflag.FindStringSubmatch(requesUrl))

	//业务
	var requesAflag string
	var requesCflag string
	if len(reAflag.FindStringSubmatch(requesUrl)) == 2 {
		requesAflag = reAflag.FindStringSubmatch(requesUrl)[1]
	}
	if len(reCflag.FindStringSubmatch(requesUrl)) == 2 {
		requesCflag = reCflag.FindStringSubmatch(requesUrl)[1]
	}

	println("==============================================================")
	//返回信息
	responseStatus := lineArr[7]
	responseSize := lineArr[8]
	responseUseTimeSec := lineArr[16]

	fmt.Printf("requesHost:|%s|\n", requesHost)
	fmt.Printf("requesClientIp:|%s|\n", requesClientIp)
	fmt.Printf("requesTime:|%s|\n", requesTime)
	fmt.Printf("requesUserAgent:|%s|\n", requesUserAgent)

	fmt.Printf("requestMethod:|%s|\n", requestMethod)
	fmt.Printf("requesUrl:|%s|\n", requesUrl)
	fmt.Printf("requesHttpVersion:|%s|\n", requesHttpVersion)

	fmt.Printf("responseStatus:|%s|\n", responseStatus)
	fmt.Printf("responseSize:|%s|\n", responseSize)
	fmt.Printf("responseUseTimeSec:|%s|\n", responseUseTimeSec)

	fmt.Printf("backendHost:|%s|\n", backendHost)
	fmt.Printf("backendPort:|%s|\n", backendPort)

	fmt.Printf("requesAflag:|%s|\n", requesAflag)
	fmt.Printf("requesCflag:|%s|\n", requesCflag)
	println("==============================================================")
}
