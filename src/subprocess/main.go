package main

import (
	"bufio"
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

func startfunc() {
	var err error
	cmd := exec.Command("python", "run2.py")
	// cmd := exec.Command("python", "run2.py")
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		log.Printf("err pip:%s\n", err)
		return
	}

	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start cmd: %v", err)
		startch <- 1
		return
	}

	var lineStr string
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		lineStr = line
		fmt.Printf("%s", lineStr)
	}
	// And when you need to wait for the command to finish:
	if err := cmd.Wait(); err != nil {
		startch <- 3
		log.Printf("Cmd returned error: %v", err)
	}
}
func restart() error {
	cmd := exec.Command("python", "run.py")

	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start cmd: %v", err)
		return err
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	fmt.Println("dir:", cmd.Dir)
	pid = cmd.Process.Pid
	fmt.Println("pid", pid)
	if err := cmd.Wait(); err != nil {
		// startch <- 3
		log.Printf("Cmd returned error: %v", err)
	}
	return nil
}
func monitorPid() {
	for {
		process, err := os.FindProcess(pid)
		if err != nil {
			err2 := restart()
			if err2 != nil {
				log.Printf("errstart:%s\n", err2)
			} else {
				log.Printf("restart success:%d\n", pid)
			}
		} else {
			log.Printf("pid[%+v] is running", process)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	startch = make(chan int, 1)
	fakechan = make(chan int, 1)
	cat()
	go startfunc()
	x := 0
	select {
	case x = <-startch:
		fmt.Println("x", x)
	}
	// if x == 3 {
	// 	panic("err start func")
	// }
	// go monitorPid()
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("subprocess pid:", process.Pid)
	select {
	case x = <-fakechan:
		fmt.Println(x)
	}

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
