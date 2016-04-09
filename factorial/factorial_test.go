package factorial

import (
    "fmt"
	"testing"
)

func TestRecursiveSmallNumber(t *testing.T) {

	fmt.Println("Test Recursive small number (10)")
    var expected int64
	fact := Factorial{Value: 10}
    go fact.ServeRecursive()
	var channel chan int64
	channel = make(chan int64)
	fact.ResultChan = channel
	res := <-channel
    expected = 3628800
	if res != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, res)
	}
}

func TestRecursiveBigNumber(t *testing.T) {

	fmt.Println("Test Recursive big number (25)")
    var expected int64
	fact := Factorial{Value: 25}
    go fact.ServeRecursive()
	var channel chan int64
	channel = make(chan int64)
	fact.ResultChan = channel
	res := <-channel
    expected = 7034535277573963776
	if res != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, res)
	}
}


func TestIterativeSmallNumber(t *testing.T) {

	fmt.Println("Test Iterative small number (10)")
    var expected int64
	fact := Factorial{Value: 10}
    go fact.ServeIterative()
	var channel chan int64
	channel = make(chan int64)
	fact.ResultChan = channel
	res := <-channel
    expected = 3628800
	if res != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, res)
	}
}

func TestIterativeBigNumber(t *testing.T) {

	fmt.Println("Test Iterative big number (25)")
    var expected int64
	fact := Factorial{Value: 25}
    go fact.ServeIterative()
	var channel chan int64
	channel = make(chan int64)
	fact.ResultChan = channel
	res := <-channel
    expected = 7034535277573963776
	if res != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, res)
	}
}

