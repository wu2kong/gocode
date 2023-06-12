package main

import (
	"log"
	"runtime"
)

func main() {
	var m = make(map[int]bool)

	for i := 0; i < 50000; i++ {
		m[i] = true
		delete(m, i)

		if i%100 == 0 {
			readMemStats()
		}
	}
}

// 打印内存占用信息
func readMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc = %v kb TotalAlloc = %v kb Sys = %v kb NumGC = %v\n", ms.Alloc/1024, ms.TotalAlloc/1024, ms.Sys/1024, ms.NumGC)
}
