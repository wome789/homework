package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

var vcount int64

func main() {
	var wg sync.WaitGroup
	s := SafeCounter{}
	/* 	for i := 0; i < 1000; i++ {
		go write(&s)
	} */

	for i := 0; i < 10; i++ {
		go add()
	}

	time.Sleep(3 * time.Second)
	fmt.Println(s.count)
	fmt.Println(vcount)
	wg.Wait()
}

/* 1.题目:编写一个程序，使用sync.Mutex来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
 */
func write(count *SafeCounter) {
	defer count.mu.Unlock()
	count.mu.Lock()
	count.count++
}

/*
	2.题目:使用原子操作(sync/atomic 包)实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

。考察点: sync.Mutex的使用、并发数据安全。
。考察点:原子操作、并发数据安全。
*/
func add() {
	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(&vcount, int64(i))
	}
}
