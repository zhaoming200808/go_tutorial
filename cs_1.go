package main

import (
	"net/http"
	"time"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/", SayHello)
	http.HandleFunc("/hello", SayHello)
	//init server
	go http.ListenAndServe("0.0.0.0:8001", nil)
	time.Sleep( 30 * time.Second )
}
