package main

import (
    "fmt"
    groupcache "github.com/golang/groupcache"
    "log"
    "net/http"
	"time"
	"strconv"
)

func main() {
    me_gc_port := ":" + "8001"
    peers := groupcache.NewHTTPPool("http://localhost" + me_gc_port)
    peers.Set("http://localhost:8002")

    helloworld := groupcache.NewGroup("helloworld_1", 1024*1024*1024*16, groupcache.GetterFunc(
        func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			value := key + me_gc_port
			log.Printf("key: %s \tvalue: %s",key,value)
			dest.SetString(value)
            return nil
    }))

    fmt.Println("GroupName: ", helloworld.Name())
	go http.ListenAndServe("localhost:" + me_gc_port, peers)

	for i:= 0 ; i <= 10 ; i++ {
		key := strconv.Itoa(int(i))
		var value string
		if err := helloworld.Get(nil, key, groupcache.StringSink(&value)); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	time.Sleep( 9000 * time.Second )
}

