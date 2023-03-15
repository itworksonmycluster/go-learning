package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var number int
	var mu sync.Mutex

	for i := 1; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			number++
			fmt.Printf("Go routine number: %d, number: %d\n", n, number)
			mu.Unlock()

		}(i)
	}

	wg.Wait()
}
