package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("ls", "-l").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)

	fmt.Printf("%v\n", err)


	cmd := exec.Command("ls", "-l")
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)


}

