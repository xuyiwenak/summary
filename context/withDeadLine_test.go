package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWithDeadLine(t *testing.T) {
	// 设置50毫秒定时时间
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel() //时间超时会自动调用

	select {
	// 设置一个超时的时间1s
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	// 超过50mscontext自动发出done，也可以修改time.After(1 * time.Millisecond),则不等到done先overslept

}
// 本质上和WithDeadline相同, 参数传递的是延迟时间
func TestWithTimeOut(t *testing.T) {
	// 设置50毫秒定时时间
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel() //时间超时会自动调用

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
func TestMutiProcessWithDeadLine(t *testing.T) {
	// 设置3秒定时时间
	d := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel() //时间超时会自动调用
	i :=0
	j :=100
	go func(ctx context.Context, cancel context.CancelFunc, i int) {
		for{
			fmt.Println("enter G1")
			select {
			case <-ctx.Done():
				fmt.Println("cancel G1 ")
				cancel()
			default:
				fmt.Printf("G1: %d\n", i)
				i++
			}
		}
	}(ctx, cancel, i)

	go func(ctx context.Context, cancel context.CancelFunc, j int) {
		for{
			fmt.Println("enter G2")
			select {
			case <-ctx.Done():
				fmt.Println("cancel G2 ")
				cancel()
			default:
				fmt.Printf("G2: %d\n", j)
				j++
			}
		}

	}(ctx, cancel, j)

	select {
	// 设置一个超时的时间1s
	case <-time.After(10 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("ready to quit")
	}
	// 超过50mscontext自动发出done，也可以修改time.After(1 * time.Millisecond),则不等到done先overslept
}
