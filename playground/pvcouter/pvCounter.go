package main

import (
	"fmt"
	"net/http"
	"time"
)

var pvCounterValue int64   // 全局的访问量统计计数器
var reportedPvValue int64  // 已上报过的Pv值，如果 reportedPvValue == pvCounterValue，那么不执行重复上报
var reqQueue chan struct{} // 使用channel模拟的日志队列

// 逻辑：从 reqQueue 通道中读取记录，进行统计计数
func pvCounter() {
	// 可以执行初始化，从服务器端拉取历史访问次数
	// pvCounterValue = redis.Get('pvCounterValue')
	fmt.Printf("%s, initial pvCounterValue is : %d.\n", time.Now().Format("2006-01-02 15:04:05"), pvCounterValue)

	// 每5秒执行上报的版本
	heartbeat := time.NewTicker(1 * time.Second)
	defer heartbeat.Stop()
	for {
		select {
		case <-reqQueue:
			pvCounterValue++
			// fmt.Printf("pvCounterValue added, now is : %d.\n", pvCounterValue)
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
	// fmt.Println("1 page viewed.")
	reqQueue <- struct{}{}
}

// http handle
func sayHello(w http.ResponseWriter, _ *http.Request) {
	defer pvAdd()
	// defer func() {
	// 	pvAdd()
	// }()

	fmt.Fprintln(w, "Hello World！")
}

func main() {
	reqQueue = make(chan struct{}, 10)
	defer close(reqQueue)
	go pvCounter()

	// http://localhost:10000/hello
	http.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":10000", nil)
}

// wrk -t10 -c30 -d 2s -T5s --latency http://localhost:10000/hello
