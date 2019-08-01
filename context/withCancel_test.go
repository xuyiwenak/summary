package context

import (
	"context"
	"fmt"
	"testing"
)

func TestWithCancel(t *testing.T)  {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done(): //只有撤销函数被调用后，才会触发
					fmt.Println("ctx done!")
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()  //调用返回的cancel方法来让 context声明周期结束

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
/*
func mapHandler(ctx context.Context) <-chan string {
	dst := make(chan string)
	dst <- names[i]
	go func() {
		for {
			select {
			case <-ctx.Done(): //只有撤销函数被调用后，才会触发
				fmt.Println("ctx done!")
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
func TestMapContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()  //调用返回的cancel方法来让 context声明周期结束
	names := []string{"lily", "yoyo", "cersei", "rose", "annei"}

	for n := range mapHandler(ctx) {
		go func() {
			fmt.Println(name)
		}()
		time.Sleep(time.Second)
	}
	runtime.GOMAXPROCS(1)
	runtime.Gosched()
}
*/