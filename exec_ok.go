package main

import "fmt"
import "bytes"
import "os/exec"
func main() {
	bey , err:= run("/tmp","pwd")
	if err != nil {
		fmt.Printf("return:%s",string(bey))
		fmt.Printf("err:%s\n",err.Error())
	}
	fmt.Printf("return:%s",string(bey))
	fmt.Println("end")
}





func run(dir string, args ...string) ([]byte, error) {
	var buf bytes.Buffer
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Stdout = &buf
	cmd.Stderr = cmd.Stdout
	err := cmd.Run()
	return buf.Bytes(), err
}




