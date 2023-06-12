package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

// 无限生成整数的goroutine，只取前5个数

// gen
func gen(ctx context.Context) chan int {
	ch := make(chan int, 100)

	go func() {
		i := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Printf("now gen: %v \n", i)
				ch <- i
				i++
				time.Sleep(time.Second)
			}
		}
	}()

	return ch
}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	for n := range gen(ctx) {
		fmt.Printf("now n: %v \n", n)
		if n == 5 {
			cancel()
			break
		}
	}

	fmt.Println("main func done")

	<-interrupt
}
