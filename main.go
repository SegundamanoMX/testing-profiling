package main

import (
	"fmt"
	"github.com/goris/testing-profiling/factorial"
	_ "os"
)

func init() {
	fmt.Println("Initialize things... ")
}

func main() {
	var factorial1 factorial.Factorial
	var factorial2 factorial.Factorial
	factorial1.Value = 6
	factorial2.Value = 6

	go factorial2.ServeRecursive()
	go factorial1.ServeIterative()

	var channel1 chan int64
	var channel2 chan int64
	channel1 = make(chan int64)
	channel2 = make(chan int64)
	factorial1.ResultChan = channel1
	factorial2.ResultChan = channel2

	res1 := <-channel1
	res2 := <-channel2

	fmt.Println("Iterative: ", res1)
	fmt.Println("Recursive: ", res2)
}
