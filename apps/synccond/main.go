package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []string
	cond  sync.Cond
	len   int
}

func main() {
	q := &Queue{
		queue: []string{},
		cond:  *sync.NewCond(&sync.Mutex{}),
		len:   5,
	}

	go func() {
		for {
			q.Inqueue("job")
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		q.Dequeue()
		time.Sleep(1 * time.Second)
	}
}

func (q *Queue) Inqueue(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) >= q.len {
		// full
		fmt.Printf("job full.\n")
		q.cond.Wait()
	}
	q.queue = append(q.queue, item)
	q.cond.Broadcast() // 全部通知
	fmt.Printf("job added.\n")
}

func (q *Queue) Dequeue() (item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) == 0 {
		// empty
		fmt.Printf("job empty.\n")
		q.cond.Wait()
	}
	item, q.queue = q.queue[0], q.queue[1:]
	q.cond.Signal() // 随机通知一个
	fmt.Printf("job done.\n")
	return
}
