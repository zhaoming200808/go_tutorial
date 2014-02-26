/**
 * http://golang.org/doc/effective_go.html#web_server
 * with groupcache (https://github.com/golang/groupcache)
 */

package main

import (
        "encoding/base64"
        "flag"
        "fmt"
        "github.com/golang/groupcache"
        "html/template"
        "io/ioutil"
        "log"
        "net/http"
)

var addr = flag.String("addr", ":8000", "http service address")
var port = flag.String("port", "8001", "groupcache http port")

var templ = template.Must(template.New("qr").Parse(templateStr))
var groupcache_group *groupcache.Group

func main() {
        flag.Parse()

        peers := groupcache.NewHTTPPool("http://localhost:" + *port)
        peers.Set("http://localhost:8001", "http://localhost:8002")

        groupcache_group = groupcache.NewGroup("QRCache", 1<<20, groupcache.GetterFunc(
                func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
                        fmt.Printf("asking for data of %s\n", key)
                        url := fmt.Sprintf("http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl=%s", key)
                        resp, err := http.Get(url)
                        if err != nil {
                                return nil
                        }
                        body, err := ioutil.ReadAll(resp.Body)
                        if err != nil {
                                return nil
                        }
                        value := base64.StdEncoding.EncodeToString(body)
                        dest.SetBytes([]byte(value))
                        return nil
                }))

        // run groupcache process in goroutine
        go http.ListenAndServe("localhost:"+*port, peers)

        http.Handle("/", http.HandlerFunc(QR))
        err := http.ListenAndServe(*addr, nil)
        if err != nil {
                log.Fatal("ListenAndServe:", err)
        }
}

func QR(w http.ResponseWriter, req *http.Request) {
        locals := make(map[string]string)
        locals["s"] = req.FormValue("s") //Blockieren
		fmt.Printf("s: %v \n",locals)
        if len(locals["s"]) > 0 {
                var value string
                if err := groupcache_group.Get(nil, locals["s"], groupcache.StringSink(&value)); err != nil {
                        fmt.Printf("err: %v\n", err)
                }
				fmt.Printf("value: |%v| \n",value)
                locals["image"] = value
        }
        templ.Execute(w, locals)
        // show groupcache stats
        fmt.Printf("####### Stats ######")
        fmt.Printf("Group Stats:\n")
        fmt.Printf("   Gets: %d\n", groupcache_group.Stats.Gets)
        fmt.Printf("   CacheHits: %d\n", groupcache_group.Stats.CacheHits)
        fmt.Printf("   PeerLoads: %d\n", groupcache_group.Stats.PeerLoads)
        fmt.Printf("   PeerErrors: %d\n", groupcache_group.Stats.PeerErrors)
        fmt.Printf("   Loads: %d\n", groupcache_group.Stats.Loads)
        fmt.Printf("   LoadsDeduped: %d\n", groupcache_group.Stats.LoadsDeduped)
        fmt.Printf("   LocalLoads: %d\n", groupcache_group.Stats.LocalLoads)
        fmt.Printf("   LocalLoadErrs: %d\n", groupcache_group.Stats.LocalLoadErrs)
        fmt.Printf("   ServerRequests: %d\n", groupcache_group.Stats.ServerRequests)
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .s}}
<img src="data:image/png;base64,{{.image}}" />
<br>
{{.s}}
<br>
<br>
{{end}}
<form action="/" method="GET">
<input maxLength=1024 size=70 name="s" value="" title="Text to QR Encode">
<input type="submit" value="Show QR">
</form>
</body>
</html>
`
