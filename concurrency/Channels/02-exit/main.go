package main

import (
	"fmt"
	"runtime"
	"time"
)

func PrintGoroutines(message string) {
	fmt.Printf("Goroutines: %d\tMessage: %s\n", runtime.NumGoroutine(), message)
}

func main() {

	exit := make(chan bool)

	for i := 1; i <= 10; i++ {
		PrintGoroutines(fmt.Sprintf("spawnning goroutine %d", i))
		go func(exit chan bool, x int) {
			for {
				select {
				case <-exit:
					fmt.Printf("Exiting goroutine %d\n", x)
					return
				default:
					fmt.Printf("Goroutine %d\n", x)
					time.Sleep(1 * time.Second)
				}
			}

		}(exit, i)
	}

	// execute goroutines for 10 sec
	time.Sleep(10 * time.Second)
	// exit goroutines
	PrintGoroutines("Exiting goroutines")
	exit <- true
	close(exit)
	// waiting 2 secs to exit goroutines
	time.Sleep(2 * time.Second)
	PrintGoroutines("End of programm")

}
