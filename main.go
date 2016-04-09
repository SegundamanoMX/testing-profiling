package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goris/testing-profiling/factorial"
	_ "os"
	"strconv"
)

var r *gin.Engine

func init() {
	fmt.Println("Initialize things... ")
	r = gin.Default()
	publics := r.Group("api/factorial")
	publics.GET("/iterative/:id", runFactorialIterative)
	publics.GET("/recursive/:id", runFactorialRecursive)
}

func main() {
	fmt.Println("Running API... ")
	r.Run()
}

func runFactorialRecursive(c *gin.Context) {
	value, err := getContextValue(c)
	if err != nil {
		c.JSON(403, gin.H{"Incorrect param": c.Params.ByName("id")})
	} else {
		var factorial factorial.Factorial
		factorial.Value = value
		go factorial.ServeRecursive()

		var channel chan int64
		channel = make(chan int64)
		factorial.ResultChan = channel
		res := <-channel
		fmt.Println("Recursive: ", res)
		c.JSON(201, gin.H{"Result": res})
	}
}

func runFactorialIterative(c *gin.Context) {
	value, err := getContextValue(c)
	if err != nil {
		c.JSON(403, gin.H{"Incorrect param": c.Params.ByName("id")})
	} else {
		var factorial factorial.Factorial
		factorial.Value = value
		go factorial.ServeIterative()

		var channel chan int64
		channel = make(chan int64)
		factorial.ResultChan = channel
		res := <-channel
		fmt.Println("Iterative: ", res)
		c.JSON(201, gin.H{"Result": res})
	}
}

func getContextValue(c *gin.Context) (int64, error) {
	var value int64
	str := c.Params.ByName("id")
	value, err := strconv.ParseInt(str, 10, 64)
	return value, err
}
