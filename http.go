package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Fprint(w, string(body))
}

func main() {
	var h Hello
	http.ListenAndServe("0.0.0.0:4000", h)
}
