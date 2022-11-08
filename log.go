package nlog

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func LogFile(path string, v ...interface{}) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println(v)
}

func LogStack(path string) {
	stack := stack()
	LogFile(path, stack)
}

func PrintStack() {
	fmt.Println(stack())
}

func LogFileWithPid(path string, v ...interface{}) {
	pid := os.Getpid()
	ppid := os.Getppid()
	pidContens, _ := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	ppidContens, _ := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", ppid))
	LogFile(path, "pid is", pid, string(pidContens))
	LogFile(path, "ppid is", ppid, string(ppidContens))
	LogFile(path, v)
}

func stack() string {
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], true)])
}
