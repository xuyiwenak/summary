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
