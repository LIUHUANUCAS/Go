package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/**
  main function
*/

var n = flag.Int("n", 1, "thread")
var s = flag.String("s", "empty", "info")

func main_28() {
	in := make(chan int)
	out := make(chan int)
	go count(in)
	go squre(in, out)

	for x := range out {
		fmt.Println(x)
	}

}

func squre(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)

}

func count(out chan<- int) {
	for i := 1; i < 10; i++ {
		out <- i
	}
	close(out)
}

func main_27() {
	in := make(chan int)
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	go func() {
		for x := range in {
			out <- x * x
		}
		close(out)
	}()

	for x := range out {
		fmt.Println(x)
	}
}

func main_25() {
	//stu = student.Student{
	//    Id : proto.Int32(1),
	//    Name : proto.String("liuhuan"),
	//}
	fmt.Printf("%d", 1)
}
func main_26() {
	str := []string{"hello", "world", "about"}
	fmt.Println(str)
	sort.Strings(str)
	fmt.Println(str)
}
func main_24() {
	str := []string{"1", "2", "3"}
	fmt.Println(str)
	str = append(str, "hello")
	fmt.Println(str)
	str = str[1:3]
	fmt.Println(str)
	str[0] = "helloworld---"
	fmt.Println(str)
}

func main_23() {
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(*n)
	fmt.Println(*s)

}
func main_22() {
	filename := os.Args[0] + ".go"
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	line20 := strings.Split(string(filedata), "\n")[:20]
	for _, line := range line20 {
		fmt.Println(line)
	}
}

func main_21() {
	filename := os.Args[0] + ".go"
	file, err := os.Open(filename)
	if err != nil {
		return
	}

	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		fmt.Println(line)
	}
}

func main_20() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		fmt.Println(line)
	}
}

func main_19() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}

func main_18() {
	ch1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "msg1"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("res1=", msg1)
	default:
		fmt.Println("go ch1")
	}

}
func main_17() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "msg1"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("res1=", msg1)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout ch1")
	}

	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- "msg2"
	}()

	select {
	case msg2 := <-ch2:
		fmt.Println("res2=", msg2)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout ch2")
	}

}
func main_16() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "chan1"
	}()

	go func() {
		ch2 <- "chan2"
	}()

	for i := 0; i < 2; i++ {

		select {
		case msg1 := <-ch1:
			fmt.Println("recv1=", msg1)
		case msg2 := <-ch2:
			fmt.Println("recv2=", msg2)
		}
	}
}

func main_15() {
	pings := make(chan string)
	pongs := make(chan string, 2)
	ping(pings, "hellworld")
	ping(pings, "this")
	pong(pings, pongs)
	pong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println(<-pongs)
}
func ping(pings chan<- string, msg string) {
	pings <- msg
	//pings <- "this"
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func worker(done chan int) {
	fmt.Println("finish compute in worker")
	done <- 1
}
func logInfo(remark map[string]interface{}) {
	remark["hello"] = "world"
	remark["1"] = 2
	tmpmap := make(map[string]interface{})
	addInfo(tmpmap)
	remark["extra"] = tmpmap

}
func addInfo(remark map[string]interface{}) {
	remark["add"] = []int{1, 2, 3, 4}
}

func main_14() {
	remarkMap := make(map[string]interface{})
	logInfo(remarkMap)
	remarkStr, _ := json.Marshal(remarkMap)
	fmt.Println(string(remarkStr))

	var i uint64 = 10086
	var s string
	s = fmt.Sprintf("%d", i)
	fmt.Println(s)

	statusProc := []int{1, 2, 3, 4}
	procArray := make([]string, 0)
	for _, v := range statusProc {
		statusStr := fmt.Sprintf("%d", v)
		procArray = append(procArray, statusStr)
	}
	statusProcStr := strings.Join(procArray, ",")
	fmt.Println(statusProcStr)
}
func main_1125() {
	str := "123"
	d, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	d += 1
	fmt.Println(d)
}
func main_1124() {
	intstr := " 1 123  34 "
	intstr = strings.TrimSpace(intstr)
	intlist := strings.Split(intstr, " ")
	fmt.Printf("list=%v", intlist)
	fmt.Printf("str=%s", strings.Join(intlist, ","))
	for _, v := range intlist {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		i, _ := strconv.Atoi(v)
		fmt.Printf("v=[%d],i+1=[%d]", v, i+1)
	}
}

func main_1123() {
	strlist := []string{"hello", "world", "Go"}
	str := strings.Join(strlist, ",")
	fmt.Println(str)
	splitlist := strings.Split(str, ",")
	fmt.Println(splitlist)
}
func main_112() {
	n := 1
	d := 0
	switch n {
	case 0:
	case 1:
		d += 11
	case 2:
		d += 20
	}
	fmt.Println("d=", d)

	var c byte
	c = 'a'
	fmt.Printf("%c", c)
	fmt.Println(c)
	t := time.Now()
	timestamp := int32(t.Unix())
	fmt.Println(timestamp)
	strlist := make([]string, 0)
	strlist = append(strlist, "hello")
	strlist = append(strlist, "word")
	str := strings.Join(strlist, ",")
	fmt.Println(str)
	splitlist := strings.Split(str, ",")
	fmt.Println(splitlist)

	user := User{"liuhuan", 25}
	fmt.Printf("%+v", user)

}

type User struct {
	name string
	age  int
}
