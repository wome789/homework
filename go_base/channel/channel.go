package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// chann1()
	channel2()
}

/*
	1.题目:编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从倒10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整教并打印出来。

。考察点:通道的基本使用、协程间通信。
*/
func chann1() {
	var wg sync.WaitGroup

	chann1 := make(chan int)
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 10; i > 0; i-- {
			chann1 <- i
		}
		close(chann1)
	}()

	go func() {
		defer wg.Done()
		for s := range chann1 {
			fmt.Println(s)
		}
	}()

	wg.Wait()

}

/*
	2.题目:实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

。考察点:通道的缓冲机制。
*/
func channel2() {
	chann1 := make(chan int, 10)

	var wg sync.WaitGroup
	wg.Add(2)

	go provider(chann1, &wg)
	go consumer(chann1, &wg)

	wg.Wait()
}

func provider(chans chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		chans <- i
	}
	close(chans)
}

func consumer(chans chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case num, ok := <-chans:
			if !ok {
				fmt.Println("Channel closed, exiting consumer")
				return
			}
			fmt.Println(num)
			// time.Sleep(100 * time.Millisecond)
		case <-time.After(5000 * time.Millisecond):
			fmt.Println("No data for 500ms, checking again")
			return
		}
	}
}
