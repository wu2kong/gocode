package main

import (
	"fmt"
	"net/http"
	"time"
)

var pvCounterValue int64
var reportedPvValue int64 // 已上报过的Pv值，如果 reportedPvValue == pvConter，那么不执行重复上报
var reqQueue chan struct{}

// 从 reqQueue 通道中读取记录，进行统计计数
func pvCounter() {
	// 无需上报的版本
	// for range reqQueue {
	// 	pvConter++
	// 	fmt.Printf("counter: %d\n", pvConter)
	// }

	// 可以执行初始化，从服务器端拉取历史访问次数
	// pvCounterValue = redis.Get('pvCounterValue')
	fmt.Printf("%s, initial pvCounterValue is : %d.\n", time.Now().Format("2006-01-02 15:04:05"), pvCounterValue)

	// 每5秒执行上报的版本
	heartbeat := time.NewTicker(5 * time.Second)
	defer heartbeat.Stop()
	for {
		select {
		case <-reqQueue:
			pvCounterValue++
			fmt.Printf("pvCounterValue added, now is : %d.\n", pvCounterValue)
		case <-heartbeat.C:
			if reportedPvValue != pvCounterValue {
				fmt.Printf("%s, pvCounterValue is : %d.\n", time.Now().Format("2006-01-02 15:04:05"), pvCounterValue)
				reportedPvValue = pvCounterValue
			} else {
				// 忽略上报
				// fmt.Printf("%s, pvCounterValue has no changed: %d.\n", time.Now().Format("2006-01-02 15:04:05"), pvCounterValue)
			}
		}
	}
}

func pvAdd() {
	fmt.Println("1 page viewed.")
	reqQueue <- struct{}{}
}

// http handle
func sayHello(w http.ResponseWriter, _ *http.Request) {
	defer func() {
		go pvAdd()
	}()

	fmt.Fprintln(w, "Hello World！")
}
