package factorial

import (
	"fmt"
)

type Factorial struct {
	Result     int64
	ResultChan chan int64
	Value      int64
}

/*
* TODO: A factorial using cache
 */
func (f *Factorial) ServeCached() {

}

func (f *Factorial) ServeIterative() {
	f.Result = 1
	for i := f.Value; i > 0; i = i - 1 {
		fmt.Println("i:", i, " val:", f.Value, " res: ", f.Result)
		f.Result = f.Result * i
	}
	f.ResultChan <- f.Result
}

func (f *Factorial) ServeRecursive() {
	f.Result = recursive(f.Value)
	f.ResultChan <- f.Result
}

func recursive(val int64) int64 {
	fmt.Println("val:", val)
	if val < 2 {
		return 1
	} else {
		return val * recursive(val-1)
	}
}
