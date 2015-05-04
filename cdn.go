package main

import (
	"fmt"
	"log"

	"errors"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Cdn struct {
	Urls         map[string][]string
	DownLoadUrls map[string][]string
	RequestCount int
	TatolSize    int
	StartTime    time.Time
	EndTime      time.Time
	TimeOut      time.Duration
	success      chan string
}

func NewCdn(Urls map[string][]string, RequestCount int) *Cdn {
	return &Cdn{
		Urls:         Urls,
		DownLoadUrls: make(map[string][]string),
		RequestCount: 20,
		TatolSize:    0,
		TimeOut:      300,
		success:      make(chan string),
	}
}

func (this *Cdn) SetTimeOut(TimeOut time.Duration) {
	this.TimeOut = TimeOut
}

func (this *Cdn) GetUrl(url string) (data []byte, err error) {
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

func (this *Cdn) GetCdnUrls() {
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
						log.Printf("Waring: parse json is not complete:%d-%d", i, this.RequestCount)
					}
					break
				}
			}
		}

	}
}

func (this *Cdn) DownLoadAllUrls() {
	this.StartTime = time.Now()
	for url, _ := range this.DownLoadUrls {
		data, err := this.GetUrl(url)
		size := strconv.Itoa(len(data))
		CheckErr(err, "size:"+size+"|download:"+url)
		this.TatolSize = this.TatolSize + len(data)
	}
	this.EndTime = time.Now()
	this.success <- "ok"
}

func (this *Cdn) CheckCdnUrls() {
	timeOut := time.After(this.TimeOut * time.Second)
	go this.DownLoadAllUrls()
	for {
		select {
		case <-this.success:
			return
		case <-timeOut:
			CheckErr(errors.New("downlod url timeout"), "downlod urls timeout")
		}
	}
}

func (this *Cdn) Show() {
	//	for cdnUrl, urls := range this.DownLoadUrls {
	//		fmt.Printf("%s: %v\n", cdnUrl, urls)
	//	}
	//	fmt.Printf("%#v\n", this.Urls)
	fmt.Printf("==================================\n")
	fmt.Printf("DownLoadUrls:%d\n", len(this.DownLoadUrls))
	if this.TatolSize < 1024 {
		fmt.Printf("TatolSize:%db\n", this.TatolSize)
	} else if this.TatolSize >= 1024*1024 {
		fmt.Printf("TatolSize:%dM\n", this.TatolSize/1024/1024)
	} else if this.TatolSize >= 1024 {
		fmt.Printf("TatolSize:%dkb\n", this.TatolSize/1024)
	}
	fmt.Printf("==================================\n")
}
func CdnJie() {
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

	cdn := NewCdn(urls, requestCount)

	cdn.SetTimeOut(300)

	cdn.GetCdnUrls()
	cdn.CheckCdnUrls()
	cdn.Show()

	fmt.Printf("==============================================================\n")
}
