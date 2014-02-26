package main

import (
    "flag"
    "html/template"
    "log"
    "net/http"
	"time"
	"regexp"
	"bytes"
	"os/exec"
	"strings"
)

var addr = flag.String("addr", ":8001", "http service address")
var templ = template.Must(template.New("qr").Parse(templateStr))
var re_num = regexp.MustCompile(`^[1-9][0-9]+$`)

func main() {
    flag.Parse()
    http.Handle("/", http.HandlerFunc(QR))
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func QR(w http.ResponseWriter, req *http.Request) {
	locals := make(map[string]string)
	locals["zone"] = req.FormValue("zone")
	zone_num := locals["zone"]

	println("url",req.Host)

}

// exec_return
const templateStr = `
<html>
<head>
<title>Update FenTian Zabbix Monitor:更新焚天监控</title>
</head>
<body>
{{if .}}
<!-- <img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" /> -->
<br>
<big>
{{.}}
</big>
<br>
<br>
{{end}}
<form action="/" name=f method="POST"><input maxLength=1024 size=70
name=zone value="" title="Text to QR Encode"><input type=submit
value="add_zone_msg" name=return>
</form>
</body>
</html>
`
