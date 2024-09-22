package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 创建第一个 Goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 1:", i)
			time.Sleep(10 * time.Millisecond) // 模拟 I/O 操作，触发线程上下文切换
		}
	}()

	// 创建第二个 Goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 2:", i)
			time.Sleep(10 * time.Millisecond) // 模拟 I/O 操作，触发线程上下文切换
		}
	}()

	// 等待所有 Goroutine 完成
	wg.Wait()
}
