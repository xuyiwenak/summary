package main

import (
	"fmt"
	"testing"
	"time"
)

func TestIncrLock(t *testing.T) {
	var a = 0
	//var locker sync.Mutex
	// 启动 1000 个协程，需要足够大
	for i := 0; i < 1000; i++ {
		go func(idx int) {
			// 如果不增加自增必然会出问题,打开注释后恢复正常
			//locker.Lock()
			//defer locker.Unlock()
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
		}(i)
	}

	// 等待 1s 结束主程序
	// 确保所有协程执行完
	time.Sleep(time.Second)
}