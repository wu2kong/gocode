package main

import (
	"sync"
	"time"
)

// func main() {
// 	ch1 := make(chan int, 2)
// 	ch1 <- 13 // fatal error: all goroutines are asleep - deadlock!
// 	n := <-ch1
// 	println(n)
// }

// func main() {
// 	ch1 := make(chan int, 2)
// 	go func() {
// 		ch1 <- 13 // 将发送操作放入一个新goroutine中执行
// 	}()
// 	n := <-ch1
// 	println(n)
// }

// func main() {
// 	ch2 := make(chan int, 1)
// 	n := <-ch2 // 由于此时ch2的缓冲区中无数据，因此对其进行接收操作将导致goroutine挂起
// 	println(n)

// 	ch3 := make(chan int, 1)
// 	ch3 <- 17 // 向ch3发送一个整型数17
// 	ch3 <- 27 // 由于此时ch3中缓冲区已满，再向ch3发送数据也将导致goroutine挂起
// }

func produce(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(1 * time.Second)
	}

	close(ch)
}

func consume(ch <-chan int) {
	for n := range ch {
		print(n)
	}
}

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}
