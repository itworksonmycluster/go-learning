package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func PrintGoRoutines(message string) {
	fmt.Printf("Routines: %d\tMessage: %s\n", runtime.NumGoroutine(), message)
}

func main() {
	PrintGoRoutines("init")
	start := time.Now()
	ch := make(chan string)
	exit := make(chan bool)
	sigChannel := make(chan os.Signal, 1)

	// SIGINT normally used by CTRL + C
	// SIGTERM program termination that can be blocked
	// SIGKILL program termination that cannot be blocked
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	PrintGoRoutines("after signal.Notify")

	//read
	go func(c <-chan string) {
		for {
			select {
			case val, ok := <-c:
				if !ok {
					fmt.Println("Channel closed. Exiting read go routine")
					return
				}
				fmt.Println(val)
			}
		}
	}(ch)

	PrintGoRoutines("read go routine")

	//write
	go func(c chan<- string, exit <-chan bool) {
		i := 1

		for {
			select {
			case <-exit:
				fmt.Println("Closing channel. Exiting write go routine")
				close(c)
				return
			default:
				c <- fmt.Sprintf("Writing into channel. Number: %d", i)
				time.Sleep(500 * time.Millisecond)
				i++
			}

		}
	}(ch, exit)

	PrintGoRoutines("write go routine")

	// need a blocker. Waiting for write function to close the channel
sigChannel:
	for {
		select {
		case sig := <-sigChannel:
			switch sig {
			case syscall.SIGINT:
				fmt.Println()
				PrintGoRoutines("SIGINT received. Exiting Go Routines. Waiting 5 secs")
				exit <- true
				time.Sleep(5 * time.Second)

				break sigChannel
			}
		default:

		}
	}

	PrintGoRoutines("after loop")

	fmt.Println(time.Since(start))
}
