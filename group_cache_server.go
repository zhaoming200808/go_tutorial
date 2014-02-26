package main

import (
	"fmt"
	"github.com/golang/groupcache"
//	"log"
	"strconv"
	"time"
	"net/http"
)

func main() {
	peers := groupcache.NewHTTPPool("http://localhost:8001")
	peers.Set("http://localhost:8002")

	//group ：名字 
	//1<<20 缓存大小 
	//groupcache.GetterFunc（ctx groupcache.Context, key string, dest groupcache.Sink)  回调函数 
    group := groupcache.NewGroup("group", 1 << 20, groupcache.GetterFunc(
        func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			// key = key
			str := "i am value ,i am key is: " + key
			if ctx == false {
				dest.SetString("")
				return nil
			}
			dest.SetString(str)
			//value := dest
//			println("---------------------------------")
//            log.Println(key)
//            log.Println(value)
            return nil
    }))

    fmt.Printf("GroupName: %s\n", group.Name())

	go http.ListenAndServe("http://localhost:8001", peers)

	var value string
	var key string
	key = "key1"
	group.Get(true,key, groupcache.StringSink(&value))
//	fmt.Printf("key   is: |%s|\n",key)

	group.Get(true,key, groupcache.StringSink(&value))
	fmt.Printf("value is: |%s|\n",value)

	for i := 1 ; i <= 10 ; i++ {
		key = "key" + strconv.Itoa(int(i))
		group.Get(true,key, groupcache.StringSink(&value))
	}
	group.Get(true,"key1", groupcache.StringSink(&value))
	fmt.Printf("value is: |%s|\n",value)

	for i := 1 ; i <= 10 ; i++ {
		key = "key" + strconv.Itoa(int(i))
		group.Get(false,key, groupcache.StringSink(&value))
	}

	group.Get(true,"key1", groupcache.StringSink(&value))
	fmt.Printf("value is: |%s|\n",value)

	fmt.Printf("####### Stats ######")
	fmt.Printf("Group Stats:\n")
	fmt.Printf("   Gets: %d\n", group.Stats.Gets)
	fmt.Printf("   CacheHits: %d\n", group.Stats.CacheHits)
	fmt.Printf("   PeerLoads: %d\n", group.Stats.PeerLoads)
	fmt.Printf("   PeerErrors: %d\n", group.Stats.PeerErrors)
	fmt.Printf("   Loads: %d\n", group.Stats.Loads)
	fmt.Printf("   LoadsDeduped: %d\n", group.Stats.LoadsDeduped)
	fmt.Printf("   LocalLoads: %d\n", group.Stats.LocalLoads)
	fmt.Printf("   LocalLoadErrs: %d\n", group.Stats.LocalLoadErrs)
	fmt.Printf("   ServerRequests: %d\n", group.Stats.ServerRequests)

	time.Sleep( 300 * time.Second )

}


