package main

import (
	// "bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var startch chan int
var fakechan chan int
var pid int
var stdout io.ReadCloser
var procstate *os.ProcessState
var seq int = 1 
func startfunc() {
	var err error
	cmd := exec.Command("python", "run2.py")
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		log.Printf("err pip:%s\n", err)
		startch <- -1 
		return
	}

	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start cmd: %v", err)
		startch <- -2
		return
	}
	pid = cmd.Process.Pid
	log.Printf("start run pid:%d\n",pid)
	startch <- seq
	seq += 1 
	// And when you need to wait for the command to finish:
	if err := cmd.Wait(); err != nil {
		log.Printf("Cmd returned error: %v", err)
	}
	startch <- seq
	seq++
	procstate = cmd.ProcessState

}
func restart() error {
	cmd := exec.Command("python", "run2.py")

	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start cmd: %v", err)
		return err
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	fmt.Println("restart dir:", cmd.Dir)
	pid = cmd.Process.Pid
	procstate = cmd.ProcessState
	fmt.Println("restart pid", pid)
	if err := cmd.Wait(); err != nil {
		log.Printf("Cmd returned error: %v", err)
	}
	startch <- seq
	seq++
	procstate = cmd.ProcessState
	return nil
}
func monitorPid() {
	x := 1
	for {
		select {
		case x = <- startch:
			go restart()
			log.Printf("restart success...%d\n",pid)
		case <- time.After(time.Second*1):
			log.Printf("pid[%+v] is running,pid:%d,%d", procstate,pid,x)
		}
	}
}

func main() {
	startch = make(chan int, 1)
	fakechan = make(chan int, 1)
	go startfunc()
	x := 0
	x = <-startch
	fmt.Println("x", x)

	if x < 0 {
		panic(fmt.Sprintf("err start func:%d",x )) 
	}
	go monitorPid()
	
	fmt.Println("subprocess pid:", pid)

	log.Printf("end start func....\n")
	select {
	case x = <-fakechan:
		fmt.Println(x)
	// case <-time.After(time.Second*10):
	// 	break
	}
	log.Printf("end\n")
}

func cat() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "run2.py")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
	

}
