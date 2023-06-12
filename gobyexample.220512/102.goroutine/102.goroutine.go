package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func job(times int) {
	i := 1
	for {
		fmt.Printf("now gen: %v*%v=%v\n", i, times, i*times)
		i++
		time.Sleep(time.Second)

		if i*times > 10 {
			println("good by times: ", times)
			return
		}
	}
}

// func main() {
// 	go func() {
// 		time.Sleep(100000 * time.Hour)
// 	}()

// 	go job(2)
// 	go job(3)

// 	select {}
// }

func main() {
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	go job(2)
	go job(3)

	<-interrupt
}
