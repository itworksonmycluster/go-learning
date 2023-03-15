package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)

	go func(c chan<- string) {
		for x := 0; x < 10; x++ {
			ch <- fmt.Sprintf("Inside\n")
		}
		close(c)
	}(ch)

	for str := range ch {
		fmt.Print(str)
	}

}
