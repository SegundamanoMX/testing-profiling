package factorial

import (
	"fmt"
)

type Factorial32 struct {
	Result     int
	ResultChan chan int
	Value      int
	Chan       bool
}

type Factorial struct {
	Result     int64
	ResultChan chan int64
	Value      int64
	Chan       bool
}

const DEBUG bool = false

/*
* TODO: A factorial using cache
 */
func (f *Factorial) ServeCached() {

}

func (f *Factorial32) ServeIterative32() {
	f.Result = 1
	for i := f.Value; i > 0; i = i - 1 {
		if DEBUG {
			fmt.Println("i:", i, " res: ", f.Result)
		}
		f.Result = f.Result * i
	}
	if f.Chan {
		f.ResultChan <- f.Result
	}
}

func (f *Factorial32) ServeRecursive32() {
	f.Result = recursive32(f.Value)
	if f.Chan {
		f.ResultChan <- f.Result
	}
}

func recursive32(val int) int {
	if DEBUG {
		fmt.Println("val:", val)
	}
	if val < 2 {
		return 1
	} else {
		return val * recursive32(val-1)
	}
}

func (f *Factorial) ServeIterative() {
	f.Result = 1
	for i := f.Value; i > 0; i = i - 1 {
		if DEBUG {
			fmt.Println("i:", i, " res: ", f.Result)
		}
		f.Result = f.Result * i
	}
	if f.Chan {
		f.ResultChan <- f.Result
	}
}

func (f *Factorial) ServeRecursive() {
	channel := make(chan int64)
	if !f.Chan {
		res := split(1, f.Value, f.Value, nil)
		f.Result = res
	} else {
		go split(1, (f.Value/2)+1, (f.Value/2)+1, channel)
		go split((f.Value/2)+2, f.Value, f.Value, channel)
		res1, res2 := <-channel, <-channel
		f.ResultChan <- res1 * res2
	}
}

func split(from, to, val int64, channel chan int64) int64 {
	res := recursive(from, to, val)
	if DEBUG {
		fmt.Println("Result half: ", res)
	}
	if channel != nil {
		channel <- res
	}
	return res
}

func recursive(from, to, val int64) int64 {
	if DEBUG {
		fmt.Println("from:", from, "to:", to, " val:", val)
	}
	if to < 2 || from > to {
		return 1
	} else {
		val = val * recursive(from, to-1, to-1)
		if DEBUG {
			fmt.Println(val, " = ", val, "*", "recursive(", from, to-1, to-1, ")")
		}
		return val
	}
}
