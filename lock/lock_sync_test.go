package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncLock(t *testing.T){
	var wg sync.WaitGroup
	seconds := []int{1,2,3,4,5}
	for i, s := range seconds {
		// 计数加 1
		wg.Add(1)
		go func(i, s int) {
			// 计数减 1
			defer wg.Done()
			fmt.Printf("goroutine%d 结束\n", i)
		}(i, s)
	}
	// 等待执行结束
	wg.Wait()
	fmt.Println("所有 goroutine 执行结束")
}