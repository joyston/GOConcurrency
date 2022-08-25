package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func ExecAtomicCounters() {
	var ops uint64

	var wgc sync.WaitGroup

	for i := 1; i <= 50; i++ {
		wgc.Add(1)
		go func() {
			for c := 1; c <= 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wgc.Done()
		}()
	}
	wgc.Wait()
	fmt.Println("ops:", ops)
}
