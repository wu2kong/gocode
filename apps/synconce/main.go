package main

import (
	"fmt"
	"sync"
)

// type Once struct {
// 	done uint32
// 	m    Mutex
// }

func main() {
	var once sync.Once
	once.Do(func() {
		fmt.Printf("abc 1")
	})
	once.Do(func() {
		fmt.Printf("abc 2")
	})
	once.Do(func() {
		fmt.Printf("abc 3")
	})
}
