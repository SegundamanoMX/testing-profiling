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
	publics.GET("/:method/:id", runFactorial)
}

func main() {
	fmt.Println("Running API... ")
	r.Run()
}

func runFactorial(c *gin.Context) {
	value, err := getContextValue(c)
	if err != nil {
		c.JSON(403, gin.H{"Incorrect param": c.Params.ByName("id")})
	}

	method := getMethodType(c)
	if method != "" {
		var factorial factorial.Factorial
		factorial.Value = value
		switch method {
		case "recursive":
			go factorial.ServeRecursive()
			break
		case "iterative":
			go factorial.ServeIterative()
			break
		default:
			c.JSON(403, gin.H{"Incorrect param": c.Params.ByName("method")})
			return
		}

		var channel chan int64
		channel = make(chan int64)
		factorial.ResultChan = channel
		res := <-channel
		fmt.Println(method, "=> ", res)
		c.JSON(201, gin.H{"Result": res})
	} else {
		c.JSON(403, gin.H{"Incorrect param": c.Params.ByName("method")})
	}
}

func getMethodType(c *gin.Context) string {
	str := c.Params.ByName("method")
	/*
	 * TODO: convert it to lowercase
	 */
	return str
}

func getContextValue(c *gin.Context) (int64, error) {
	var value int64
	str := c.Params.ByName("id")
	value, err := strconv.ParseInt(str, 10, 64)
	return value, err
}
