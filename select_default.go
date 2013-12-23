package main

import (
    "fmt"
    "time"
)

func main() {
    tick := time.Tick(200 * time.Millisecond)
    boom := time.After(1000 * time.Millisecond)
    for {
        select {
        case <-tick:
            fmt.Println("200 tick.")
        case <-boom:
            fmt.Println("1000 BOOM!")
            return
        default:
            fmt.Println("50    .")
            time.Sleep(50 * time.Millisecond)
        }
    }
}

