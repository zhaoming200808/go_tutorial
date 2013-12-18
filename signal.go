package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 0)
	signal.Notify(c, os.Interrupt, os.Kill,syscall.SIGTERM,syscall.SIGHUP)

	// Block until a signal is received.
	for {
		time.Sleep(1 * time.Second)
		println("= , =")
	}
	s := <-c
		fmt.Println("Got signal:", s)
		return
}

