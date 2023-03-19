package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	exit := make(chan bool)

	//read
	go func(c <-chan string) {
		for {
			select {
			case val, ok := <-c:
				if !ok {
					fmt.Println("channel closed")
					exit <- true
					return
				}
				fmt.Println(val)
			}
		}
	}(ch)

	//write
	go func(c chan<- string) {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("Message id: %d", i)
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("closing the channel")
		close(c)
	}(ch)

	// need a blocker. Waiting for write function to close the channel
exitChannel:
	for {
		select {
		case <-exit:
			break exitChannel
		default:
		}
	}

	fmt.Println(time.Since(start))
}
