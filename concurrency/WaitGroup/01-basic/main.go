package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Printf("number de go routines: %d\n", runtime.NumGoroutine())

	var wg sync.WaitGroup

	for i := 1; i < 6; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Go routine number: %d, number de go routines: %d\n", n, runtime.NumGoroutine())
			time.Sleep(2 * time.Second)
		}(i)
	}

	wg.Wait()

	fmt.Printf("number de go routines: %d\n", runtime.NumGoroutine())

}
