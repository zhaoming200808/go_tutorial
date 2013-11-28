package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
//	"strings"
)

func main() {
//	cmd := exec.Command("tr", "a-z", "A-Z")
//	cmd := exec.Command("ls","-l")
	cmd := exec.Command("ls","-l")
//	cmd.Stdin = strings.NewReader("some input")
/*
    err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Waiting for command to finish...")
    err = cmd.Wait()
    log.Printf("Command finished with error: %v", err)
*/
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %s\n", out.String())
	fmt.Printf("err: %v\n", err)
}



