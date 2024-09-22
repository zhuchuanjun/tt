package main

import (
	"fmt"
	"time"
)

func main() {
	scheduler := NewScheduler(3)

	// 添加任务
	for i := 1; i <= 5; i++ {
		id := i
		scheduler.AddTask(id, func() error {
			fmt.Printf("Task %d is running\n", id)
			time.Sleep(2 * time.Second)
			fmt.Printf("Task %d is completed\n", id)
			return nil
		})
	}

	// 启动调度器
	go scheduler.Run()

	// 取消任务2和任务4
	//time.Sleep(1 * time.Second) // 确保调度器已经启动
	//scheduler.CancelTask(2)
	scheduler.CancelTask(4)

	// 查询任务状态
	time.Sleep(5 * time.Second) // 等待任务执行完成
	for i := 1; i <= 5; i++ {
		status := scheduler.GetTaskStatus(i)
		fmt.Printf("Task %d status: %s\n", i, status)
	}
}
