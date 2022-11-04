package nlog

import (
	"fmt"
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

func stack() string {
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], true)])
}
