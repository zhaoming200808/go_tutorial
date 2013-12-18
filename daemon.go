package main

import (
    "fmt"
    "log"
    "os"
    "runtime"
    "syscall"
    "time"
)

func daemon(nochdir, noclose int) int {
    var ret, ret2 uintptr
    var err syscall.Errno

    darwin := runtime.GOOS == "darwin"

    // already a daemon
    if syscall.Getppid() == 1 {
        return 0
    }

    // fork off the parent process
    ret, ret2, err = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
    if err != 0 {
        return -1
    }

    // failure
    if ret2 < 0 {
        os.Exit(-1)
    }

    // handle exception for darwin
    if darwin && ret2 == 1 {
        ret = 0
    }

    // if we got a good PID, then we call exit the parent process.
    if ret > 0 {
        os.Exit(0)
    }

    /* Change the file mode mask */
    _ = syscall.Umask(0)

    // create a new SID for the child process
    s_ret, s_errno := syscall.Setsid()
    if s_errno != nil {
        log.Printf("Error: syscall.Setsid errno: %d", s_errno)
    }
    if s_ret < 0 {
        return -1
    }
	println(nochdir )
    if nochdir == 0 {
        os.Chdir("/")
    }

	println(nochdir )
    if noclose == 0 {
        f, e := os.OpenFile("/dev/null", os.O_RDWR, 0)
        if e == nil {
            fd := f.Fd()
            syscall.Dup2(int(fd), int(os.Stdin.Fd()))
            syscall.Dup2(int(fd), int(os.Stdout.Fd()))
            syscall.Dup2(int(fd), int(os.Stderr.Fd()))
        }
    }

    return 0
}

func main() {
	time.Sleep(1 * time.Second)
	r := daemon(0, 1)
	println("daemon: ",r)
	t()
	t()
	t()
	t()
	t()
//	for i:=1; i<100 ; i++ {
//        fmt.Println("hello ",i)
//        time.Sleep(20 * time.Second)
//   }

}

func t(){
	i := 1
	fmt.Println("hello ",i)
	i = i+1
	time.Sleep(1 * time.Second)
}


