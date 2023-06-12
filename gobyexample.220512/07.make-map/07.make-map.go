package main

import "fmt"

func main() {
	/// make slice,map,channel

	// slice
	sliceA := make([]string, 3)
	sliceA[0] = "wuk"
	fmt.Printf("%v\n", sliceA)

	// map
	mapA := make(map[string]int)
	mapA["lihao"] = 18
	fmt.Printf("%v\n", mapA)
	fmt.Printf("mapA type: %T  value: %v\n", mapA, mapA)

	// channel
	channelA := make(chan int, 3)
	fmt.Printf("%v\n", channelA)

	/// new 任意类型，产生一个指针, 返回一个指针

	// slice
	newSliceB := new([]string)
	newSliceB = &sliceA
	fmt.Printf("%v\n", newSliceB)

	// map
	newMapB := new(map[string]int)
	newMapB = &mapA
	mapA["lisi"] = 20
	fmt.Printf("%v\n", newMapB)

	// channel
	newChannelB := new(chan int)
	fmt.Printf("%v\n", newChannelB)
	fmt.Printf("newChannelB type: %T  value: %v\n", newChannelB, *newChannelB)

	id := new(int)
	name := new(string)
	flag := new(bool)
	fmt.Printf("id type: %T  value: %v\n", id, *id)
	fmt.Printf("name type: %T  value: %v\n", name, *name)
	fmt.Printf("flag type: %T  value: %v\n", flag, *flag)
}
