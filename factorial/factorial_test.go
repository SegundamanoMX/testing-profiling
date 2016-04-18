package factorial

import (
	"fmt"
	"testing"
)

func TestIterative32(t *testing.T) {

	fmt.Println("Test Iterative 32bits (Input: 20)")
	var expected int
	fact := Factorial32{Value: 20, Chan: true, ResultChan: make(chan int)}
	go fact.ServeIterative32()
	res := <-fact.ResultChan
	expected = 2432902008176640000
	if res != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, res)
	}
}

func TestRecursive32(t *testing.T) {

	fmt.Println("Test Recursive 32bits (Input: 20)")
	var expected int
	fact := Factorial32{Value: 20, Chan: true, ResultChan: make(chan int)}
	go fact.ServeRecursive32()
	res := <-fact.ResultChan
	expected = 2432902008176640000
	if res != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, res)
	}
}

func TestIterative64(t *testing.T) {

	fmt.Println("Test Iterative 64bits (Input: 20)")
	var expected int64
	fact := Factorial{Value: 20, Chan: true, ResultChan: make(chan int64)}
	go fact.ServeIterative()
	res := <-fact.ResultChan
	expected = 2432902008176640000
	if res != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, res)
	}
}

func TestRecursive64(t *testing.T) {

	fmt.Println("Test Recursive 64bits (Input: 20)")
	var expected int64
	fact := Factorial{Value: 20, Chan: false, ResultChan: make(chan int64)}
	fact.ServeRecursive()
	expected = 2432902008176640000
	if fact.Result != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, fact.Result)
	}
}

func TestRecursive64Channel(t *testing.T) {

	fmt.Println("Test Recursive 64bits (Input: 20)")
	var expected int64
	fact := Factorial{Value: 20, Chan: true, ResultChan: make(chan int64)}
	go fact.ServeRecursive()
	res := <-fact.ResultChan
	expected = 2432902008176640000
	if res != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, res)
	}
}

func BenchmarkRecursive32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fact := Factorial32{Value: 20, Chan: false}
		fact.ServeRecursive32()
	}
}

func BenchmarkIterative32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fact := Factorial32{Value: 20, Chan: false}
		fact.ServeIterative32()
	}
}

func BenchmarkRecursive32GoRoutine(b *testing.B) {
	fact := Factorial32{Value: 20, Chan: true, ResultChan: make(chan int)}
	for n := 0; n < b.N; n++ {
		go fact.ServeRecursive32()
		_ = <-fact.ResultChan
	}
}

func BenchmarkIterative32GoRoutine(b *testing.B) {
	fact := Factorial32{Value: 20, Chan: true, ResultChan: make(chan int)}
	for n := 0; n < b.N; n++ {
		go fact.ServeIterative32()
		_ = <-fact.ResultChan
	}
}

func BenchmarkRecursive(b *testing.B) {
	fact := Factorial{Value: 20, Chan: false}
	for n := 0; n < b.N; n++ {
		fact.ServeRecursive()
	}
}

func BenchmarkIterative(b *testing.B) {
	fact := Factorial{Value: 20, Chan: false}
	for n := 0; n < b.N; n++ {
		fact.ServeIterative()
	}
}

func BenchmarkRecursiveGoRoutine(b *testing.B) {
	fact := Factorial{Value: 20, Chan: true, ResultChan: make(chan int64)}
	for n := 0; n < b.N; n++ {
		go fact.ServeRecursive()
		_ = <-fact.ResultChan
	}
}

func BenchmarkIterativeGoRoutine(b *testing.B) {
	fact := Factorial{Value: 20, Chan: true, ResultChan: make(chan int64)}
	for n := 0; n < b.N; n++ {
		go fact.ServeIterative()
		_ = <-fact.ResultChan
	}
}
