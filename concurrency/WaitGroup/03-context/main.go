package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func dbQuery(ctx context.Context, wg *sync.WaitGroup) {
	for {
		time.Sleep(300 * time.Millisecond)
		n := rand.Int31n(50)
		if n == 5 {
			fmt.Printf("go routine %d found 5\n", ctx.Value("goroutine"))
			break
		}
		select {
		case <-ctx.Done():
			fmt.Printf("go routine %d did not find 5\n", ctx.Value("goroutine"))
			wg.Done()
		default:
		}
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 6; i++ {
		ctx, cancel := context.WithTimeout(context.WithValue(context.TODO(), "goroutine", i), 10*time.Second)
		defer cancel()
		wg.Add(1)
		go dbQuery(ctx, &wg)

	}

	wg.Wait()

}
