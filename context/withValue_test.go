package context

import (
	"context"
	"fmt"
	"testing"
)

func TestWithValue(t *testing.T) {
	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	k1 := favContextKey("Chinese")
	ctx := context.WithValue(context.Background(), k, "Go")
	ctx1 := context.WithValue(ctx, k1, "Go1")

	f(ctx1, k1)
	f(ctx1, k)
}