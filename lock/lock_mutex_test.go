package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)
// 互斥锁是用来保证在任一时刻， 只能有一个协程访问某个对象。 Mutex的初始值为解锁的状态。
// 通常作为其他结构体的你名字段使用， 并且可以安全的在多个例程中并行使用。
func TestMutexLock(t *testing.T)  {
	ch := make(chan struct{}, 2)

	var l sync.Mutex
	go func() {
		fmt.Println("groutine1: 进入go func... time:"+ time.Now().Format(time.Stamp))
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine1: 我会锁定大概 2s time:" + time.Now().Format(time.Stamp))
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1: 我解锁了，你们去抢吧 time:"+ time.Now().Format(time.Stamp))
		ch <- struct{}{}
	}()
	go func() {
		fmt.Println("groutine2: 进入go func... time:"+ time.Now().Format(time.Stamp))
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine2: 我会锁定大概 3s time:"+ time.Now().Format(time.Stamp))
		time.Sleep(time.Second * 3)
		fmt.Println("goroutine2: 释放锁了 time:"+ time.Now().Format(time.Stamp))
		ch <- struct{}{}
	}()

	// 等待 goroutine 执行结束
	for i := 0; i < 2; i++ {
		<-ch
	}
}