package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"errors"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"strconv"
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

type Api struct {
	Urls         map[string][]string
	DownLoadUrls map[string][]string
	RequestCount int
	TatolSize    int
	StartTime    time.Time
	EndTime      time.Time
}

func NewApi(Urls map[string][]string, RequestCount int) *Api {
	return &Api{
		Urls:         Urls,
		DownLoadUrls: make(map[string][]string),
		RequestCount: 20,
		TatolSize:    0,
	}
}

func (this *Api) GetUrl(url string) (data []byte, err error) {
	var maxTry int = 3
	var errSleep time.Duration = 2

	for i := 1; i <= maxTry; i++ {
		resp, err := http.Get(url)
		if err != nil {
			time.Sleep(errSleep * time.Second)
		}
		if err != nil {
			continue
		}

		defer resp.Body.Close()
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		} else {
			break
		}
	}
	return
}

func (this *Api) GetCdnUrls() {
	for url, flags := range this.Urls {
		//fmt.Printf("%s,%s\n", url, flags)
		data, err := this.GetUrl(url)
		CheckErr(err, "get url:"+url)
		js, err := simplejson.NewJson(data)
		CheckErr(err, "parse json:"+url)

		for _, flag := range flags {
			//fmt.Printf("%s %s\n", url, flag)
			for i, j := 0, 100000; i < j; i++ {
				url, err := js.Get("list").GetIndex(i).Get(flag).String()
				if err == nil && url != "" {
					if len(this.DownLoadUrls[url]) == 0 {
						this.DownLoadUrls[url] = []string{url + "|" + flag}
					} else {
						this.DownLoadUrls[url] = append(this.DownLoadUrls[url], url+"|"+flag)
					}
				} else {
					if i == 0 {
						err := errors.New("no find flag")
						CheckErr(err, "url:"+url)
					} else if i != this.RequestCount {
						log.Printf("Waring: parse json is not complete")
					}
					break
				}
			}
		}

	}
}

func (this *Api) CheckCdnUrls() {
	for url, _ := range this.DownLoadUrls {
		data, err := this.GetUrl(url)
		size := strconv.Itoa(len(data))
		CheckErr(err, "size:"+size+"|download:"+url)
		this.TatolSize = this.TatolSize + len(data)
	}
}

func (this *Api) Show() {
	//	for cdnUrl, urls := range this.DownLoadUrls {
	//		fmt.Printf("%s: %v\n", cdnUrl, urls)
	//	}
	//	fmt.Printf("%#v\n", this.Urls)
	fmt.Printf("DownLoadUrls:%d\n", len(this.DownLoadUrls))
	if this.TatolSize < 1024 {
		fmt.Printf("TatolSize:%db\n", this.TatolSize)
	} else if this.TatolSize >= 1024*1024 {
		fmt.Printf("TatolSize:%dM\n", this.TatolSize/1024/1024)
	} else if this.TatolSize >= 1024 {
		fmt.Printf("TatolSize:%dkb\n", this.TatolSize/1024)
	}
}
func main() {
	fmt.Printf("==============================================================\n")

	requestCount := 20
	//map[string][]string
	data_url := "http://api.budejie.com/api/api_open.php?c=data&a=list&per=20"
	voice_url := "http://api.budejie.com/api/api_open.php?c=voice&a=list&per=20"
	video_url := "http://api.budejie.com/api/api_open.php?c=video&a=list&per=20"

	data_flags := []string{"image0", "cdn_img"}
	voice_flags := []string{"image0", "cdn_img", "voiceuri"}
	video_flags := []string{"image0", "cdn_img", "videouri"}

	urls := make(map[string][]string)
	urls[data_url] = data_flags
	urls[voice_url] = voice_flags
	urls[video_url] = video_flags

	api := NewApi(urls, requestCount)

	api.GetCdnUrls()
	api.CheckCdnUrls()
	api.Show()

	fmt.Printf("==============================================================\n")
}
