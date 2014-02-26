package main

import (
    "fmt"
    groupcache "github.com/golang/groupcache"
    "log"
    "net/http"
//    "os"
    "strings"
	"time"
	"strconv"
)

func main() {
    // Usage: ./test_groupcache port
    // me := ":" + os.Args[1]
    me := ":" + "8001"
    peers := groupcache.NewHTTPPool("http://localhost" + me)
    peers.Set("http://localhost:8081", "http://localhost:8082")

    helloworld := groupcache.NewGroup("helloworld", 1024*1024*1024*16, groupcache.GetterFunc(
        func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			value := key + me
			log.Println(" key : " + key)
			log.Println("value: " + value)
			dest.SetString(value)
            return nil
    }))

    fmt.Println("GroupName: ", helloworld.Name())
	for i := 1 ; i<= 10 ; i++{
		key := strconv.Itoa(int(i))
		var value []byte
		helloworld.Get(nil, key, groupcache.AllocatingByteSliceSink(&value))
	}

	go http.ListenAndServe("localhost:" + me, peers)

	time.Sleep( 9000 * time.Second )
    http.HandleFunc("/xbox/", func(w http.ResponseWriter, r *http.Request) {
        parts := strings.SplitN(r.URL.Path[len("/xbox/"):], "/", 1)
        if len(parts) != 1 {
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        }
        var data []byte
        helloworld.Get(nil, parts[0], groupcache.AllocatingByteSliceSink(&data))
        w.Write(data)
		println("----------------------------------------------")
        log.Println("Gets: ", helloworld.Stats.Gets.String())
        log.Println("Load: ", helloworld.Stats.Loads.String())
        log.Println("LocalLoad: ", helloworld.Stats.LocalLoads.String())
        log.Println("PeerError: ", helloworld.Stats.PeerErrors.String())
        log.Println("PeerLoad: ", helloworld.Stats.PeerLoads.String())
    })

    http.ListenAndServe(me, nil)
}

