package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "a log:", log.Lshortfile)
	// logger.SetFlags(log.Llongfile)
	fmt.Println(logger.Flags())

	logger.Println("hello logger")

	logger.SetPrefix("debug:")
	fmt.Println(logger.Prefix()) // debug:

	logger.Print("hello, log file")

	err := logger.Output(2, "world")
	fmt.Println(err)

	logger.Printf("%v\t%v", "hello", "go")

	// logger.Fatal("fatal log")
	// logger.Panic("Panic log")

	// 使用预定义日志类型
	fmt.Println(log.Flags())         // 3
	fmt.Println(log.Prefix())        // ''
	log.Print("hello predefine log") // 2017/03/04 16:34:27 hello predefine log
}
