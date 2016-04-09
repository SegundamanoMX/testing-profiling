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
			fmt.Println("i:", i, " val:", f.Value, " res: ", f.Result)
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
	f.Result = recursive(f.Value)
	if f.Chan {
		f.ResultChan <- f.Result
	}
}

func recursive(val int64) int64 {
	if DEBUG {
		fmt.Println("val:", val)
	}
	if val < 2 {
		return 1
	} else {
		return val * recursive(val-1)
	}
}
