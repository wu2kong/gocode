package main

import "fmt"

func main() {
	s := make([]int, 0)

	oldCap := cap(s)

	for i := 1; i <= 2048; i++ {
		s = append(s, i)

		newCap := cap(s)

		if newCap != oldCap {

			fmt.Printf("[%d -> %4d] after append 第%-4d个元素 | oldcap = %-4d | newcap = %-4d\n", 0, i-1, i, oldCap, newCap)
			oldCap = newCap
		}
	}
}
