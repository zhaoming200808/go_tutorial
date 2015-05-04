package main

import (
	"github.com/robfig/fsnotify"
	"log"
	"strings"
	"time"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				arr_ev_str := strings.SplitN(ev.String(), " ", -1)
				eventName := arr_ev_str[len(arr_ev_str)-1]
				if eventName == "MODIFY" {
					log.Println("eventName:", eventName)
					//read file 1line or buff
				} else if eventName == "RENAME" {
					log.Println("eventName:", eventName)
					// sleep and reopen file
				} else if eventName == "DELETE" {
					log.Println("eventName:", eventName)
					// sleep and reopen file
				} else {
					log.Println("WARING: eventName not find")
					log.Println("eventName:", eventName)
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("/tmp/foo")
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(100000000000000)
	}
}

type Work struct {
	FilePath    string
	StartTime   string
	ModifyCount int
	DeleteCount int
	RenameCount int
	IsRuning    bool
	FileOpenFd  int
	FileWatcher *fsnotify.Watcher
}

func (this Work) Start() bool {
	if this.IsRuning != true {
		this.IsRuning = true
	}
	return true
}

func (this Work) Stop() bool {
	if this.IsRuning != false {
		this.IsRuning = false
	}
	return true
}

func (this Work) Close() bool {
	this.Stop()
	this.FileOpenFd=0
	this.FileWatcher=nil
	return true
}

func NewWork(filePath string) *Work {
	return &Work{
		FilePath:    filePath,
		StartTime:   "",
		ModifyCount: 0,
		DeleteCount: 0,
		RenameCount: 0,
		IsRuning:    false,
	}
}

