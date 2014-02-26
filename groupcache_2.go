package main

import (
    "fmt"
    groupcache "github.com/golang/groupcache"
    "log"
    "net/http"
//    "os"
    "strings"
)

func main() {
	http_port := ":8000"
    me_gc_port := ":" + "8002"
    peers := groupcache.NewHTTPPool("http://localhost" + me_gc_port)
    peers.Set("http://localhost:8001")

    helloworld := groupcache.NewGroup("helloworld", 1024*1024*1024*16, groupcache.GetterFunc(
        func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			value := key + me_gc_port
			log.Println("key: "+key)
			dest.SetString(value)
            return nil
    }))

    fmt.Println("GroupName: ", helloworld.Name())
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
	http.ListenAndServe("localhost:" + me_gc_port, peers)
	http.ListenAndServe(http_port, nil)
}

