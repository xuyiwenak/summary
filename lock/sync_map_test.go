package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
	"sync"
)
const N int = 1000

// N太小时不会（比如10），因机器而异
// fatal error: concurrent map read and map write
func TestMap(t *testing.T) {
	m := make(map[string]int)

	go func() {
		for i := 0; i < N; i++ {
			m[strconv.Itoa(i)] = i // write
		}
	}()

	go func() {
		for i := 0; i < N; i++ {
			fmt.Println(i, m[strconv.Itoa(i)]) // read
		}
	}()

	time.Sleep(time.Second * 5)
}
// 使用sync.map线程安全
func TestSyncMap(t *testing.T)  {
	var m sync.Map

	go func() {
		for i := 0; i < N; i++ {
			m.Store(strconv.Itoa(i), i) // 写
		}
	}()

	go func() {
		for i := 0; i < N; i++ {
			v, _ := m.Load(strconv.Itoa(i)) // 读
			fmt.Println(i, v)
		}
	}()

	time.Sleep(time.Second * 5)
}