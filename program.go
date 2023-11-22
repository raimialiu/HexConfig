package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("goroutine test")
	fmt.Println("this should run first")
	functionA("olatunde")
	//time.Sleep(time.Second * 2000)
	fmt.Println("this should run last")
}

func functionA(name string) {
	defer cleanUpCode()
	panic("panicking this function to test bubbling up....")
	printName(name)
	fmt.Println("after panicking.....")
}

func printName(name string) {
	time.Sleep(time.Second * 2)
	fmt.Println("my name is " + name)
}

func cleanUpCode() {
	fmt.Println("clean up coding.....")
	time.Sleep(time.Second * 2)

	recoveryFunc := func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(string); ok {
				fmt.Println("catching....." + err)
			} else if strErr, ok := arg.(error); ok {
				fmt.Println(strErr.Error())
			} else {
				fmt.Println("recovery from error")
			}
		}
	}

	recoveryFunc()
}
