package main

import "fmt"

func main() {
	var sl = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}

	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl[i]); j++ {
			if sl[i][j] == 13 {
				fmt.Printf("found 13 at [%d, %d]\n", i, j)
				break
			}
		}
	}

	print(1, 2, 3)
}
