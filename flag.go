package main

import (
	"flag"
	"fmt"
)

func main() {
	serviceName := flag.String("n", "mysql", "defaule serviceName")
	serviceHost := flag.String("h", "127.0.0.1", "defaule host or ip")
	servicePort := flag.Int("p", 3306, "defaule port")

	flag.Parse()
	fmt.Printf("serviceName: %s\nserviceHost: %s\nservicePort: %d\n", *serviceName, *serviceHost, *servicePort)
}
