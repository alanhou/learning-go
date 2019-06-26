package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a * atomicInt) incement() {
	fmt.Println("safe increment")
	func () {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atomicInt
	a.incement()
	go func() {
		a.incement()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
