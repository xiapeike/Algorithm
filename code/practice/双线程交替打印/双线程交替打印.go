package main

import (
	"fmt"
	"sync"
)

func printEven(wg *sync.WaitGroup, evenCh chan int, oddCh chan int) {
	defer wg.Done()

	for i := 2; i <= 10; i += 2 {
		<-evenCh // 等待接收 evenCh 的值
		fmt.Println(i)
		oddCh <- 1 // 发送一个到 oddCh
	}
}

func printOdd(wg *sync.WaitGroup, evenCh chan int, oddCh chan int) {
	defer wg.Done()

	for i := 1; i <= 9; i += 2 {
		<-oddCh // 等待接 oddCh 的值
		fmt.Println(i)
		evenCh <- 1 // 发送一个值到 evenCh
	}
}

func main() {
	var wg sync.WaitGroup
	evenCh := make(chan int)
	oddCh := make(chan int)

	wg.Add(2)

	go printEven(&wg, evenCh, oddCh)
	go printOdd(&wg, evenCh, oddCh)

	// 启动两个 goroutine 后，先发送一个值到 evenCh以触发打印偶数的 goroutine 开始执行
	oddCh <- 1

	wg.Wait()

}
