package main

import (
	"fmt"
	"os"
	"time"
)

func fdCopyRead(fd *os.File) {
	// 打印描述符地址
	fmt.Println(fd)

	_, err := fd.Seek(0, 0)
	b1 := make([]byte, 13)
	n1, err := fd.Read(b1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
}

func main() {
	filepath := "../../.gitignore"

	fd, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	defer fd.Close()

	// 打印描述符地址
	fmt.Println(fd)

	b1 := make([]byte, 10)
	n1, err := fd.Read(b1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	go fdCopyRead(fd)

	time.Sleep(5 * time.Second)
}

// content, err := ioutil.ReadAll(fd)
// fmt.Println(content)
