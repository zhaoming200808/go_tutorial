package main

import "fmt"
import "os"
import "bytes"
import "os/exec"
import "time"
import "errors"
import "runtime"

func main() {
	for i, n := 1, 1000000; n > i; i++ {
		fmt.Printf("index: %d \truntime.NumGoroutine: %d\n", i, runtime.NumGoroutine())
		_, err := run(".", 100, "ls -al")
		//_, err := t_exec(".", "ls")
		//err := t_chan()
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(10000000000 * time.Millisecond)
}

func t_chan() error {
	quit := make(chan bool)
	value := make(chan string)
	defer close(quit)
	defer close(value)

	go func(quit chan bool) {
		time.Sleep(1000 * time.Millisecond)
		_, ok := <-quit
		if ok {
			println("quit is ok!")
			quit <- true
		} else {
			println("quit is not ok!")
		}
	}(quit)

	go func(value chan string) {
		value <- "ok"

	}(value)

	select {
	case <-value:
		return nil
	case <-quit:
		return nil
	}
}

func t_exec(dir string, args ...string) (v []byte, err error) {
	args = append(args, "/bin/bash", "-c")
	var buf bytes.Buffer
	var cmd *exec.Cmd

	cmd = exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Stdout = &buf
	//cmd.Stderr = cmd.Stdout
	err = cmd.Run()
	//cmd.ProcessState.Exited()
	v = buf.Bytes()

	return v, err
}

func run(dir string, t time.Duration, c ...string) (v []byte, err error) {
	//check args
	args := make([]string, 0, 10)
	args = append(args, "/bin/bash", "-c")
	for _, v := range c {
		args = append(args, v)
	}

	quit := make(chan bool)
	value := make(chan string)
	defer close(quit)
	defer close(value)

	var buf bytes.Buffer
	var cmd *exec.Cmd
	go func(quit chan bool) {
		time.Sleep(t * time.Millisecond)
		_, ok := <-quit
		if ok {
			quit <- true
		}
	}(quit)

	go func(value chan string) {
		cmd = exec.Command(args[0], args[1:]...)
		cmd.Dir = dir
		cmd.Stdout = &buf
		cmd.Stderr = cmd.Stdout
		err = cmd.Run()
		cmd.ProcessState.Exited()
		v = buf.Bytes()
		value <- "ok"
	}(value)
	select {
	case <-value:
		cmd.ProcessState.Exited()
		return v, err
	case <-quit:
		if cmd != nil {
			fmt.Printf("%v\n", cmd.Process.Kill())
			return nil, errors.New("run cmd time out")
		}
		return nil, nil
	}
	return
}

func t_run(dir string, t time.Duration, c ...string) (v []byte, err error) {
	//check dir
	file, err := os.Stat(dir)
	if err != nil || file.IsDir() != true {
		fmt.Println("%s", err.Error())
		return nil, errors.New("path not find: " + dir)
	}
	//check args
	args := make([]string, 0, 10)
	args = append(args, "/bin/bash", "-c")
	for _, v := range c {
		args = append(args, v)
	}

	quit := make(chan bool)
	value := make(chan string)
	var buf bytes.Buffer
	var cmd *exec.Cmd
	go func(quit chan bool) {
		time.Sleep(t * time.Millisecond)
		quit <- true
	}(quit)

	go func(value chan string) {
		cmd = exec.Command(args[0], args[1:]...)
		cmd.Dir = dir
		cmd.Stdout = &buf
		cmd.Stderr = cmd.Stdout
		err = cmd.Run()
		cmd.ProcessState.Exited()
		v = buf.Bytes()
		value <- "ok"
	}(value)
	select {
	case <-value:
		cmd.ProcessState.Exited()
		return v, err
	case <-quit:
		if cmd != nil {
			fmt.Printf("%v\n", cmd.Process.Kill())
			return nil, errors.New("run cmd time out")
		}
		return nil, nil
	}
}
