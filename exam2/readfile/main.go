package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"
)

func main() {
	var f = "/var/logs/app.log"
	file, err := os.OpenFile(f, os.O_RDWR, os.ModeExclusive)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 调用系统调用加锁
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		panic(err)
	}
	defer syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
	// 读取文件内容
	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", all)
	time.Sleep(time.Second * 10) //模拟耗时操作
}
