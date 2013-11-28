package main

import (
    "fmt"
    "time"
	"runtime"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(1000 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
	runtime.GOMAXPROCS(200)
    go say("--world @go")
    go say("**world @go")
    go say("~~world @go")
    say("++hello @main++")
}




