package main

import (
	"fmt"
	"time"
)

func main() {

	exit := make(chan bool)

	for i := 1; i <= 10; i++ {
		go func(exit chan bool, x int) {
			for {
				select {
				case <-exit:
					return
				default:
					fmt.Printf("Go routine %d\n", x)
					time.Sleep(1 * time.Second)
				}
			}

		}(exit, i)
	}

	time.Sleep(10 * time.Second)
	exit <- true
	close(exit)

}
