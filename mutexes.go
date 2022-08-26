package main

import (
	"fmt"
	"sync"
)

type container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func ExecMutex() {
	c := container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var mwg sync.WaitGroup
	doIncrement := func(name string, n int) {

		for i := 0; i < n; i++ {
			c.inc(name)
		}
		mwg.Done()
	}

	mwg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	mwg.Wait()

	fmt.Println(c.counters)
}
