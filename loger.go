package main

import (
	"fmt"
	"log"
	"os"
)



func main(){
	var logger *log.Logger
	logfile,err := os.OpenFile("/var/log/admin.log",os.O_CREATE|os.O_APPEND|os.O_RDWR,0660);
	if err!=nil {
		fmt.Printf("%s\r\n",err.Error());
		logger = log.New(os.Stderr, "admin: ", log.Ldate | log.Ltime | log.Lshortfile)
	}else{
		logger = log.New(logfile,"",log.Ldate|log.Ltime|log.Llongfile);
	}
	defer logfile.Close();

//	logger = log.New(os.Stderr, "", log.Ldate | log.Ltime | log.Lshortfile)
	logger.Println("hello");
	logger.Println("oh....");
	logger.Println("oh....");
}


