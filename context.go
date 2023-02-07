package main

import (
	"context"
	"fmt"
)

// func main() {
// 	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
// 	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	go handle(ctx, 500*time.Millisecond)
// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("main", ctx.Err())
// 	}
// }

// func handle(ctx context.Context, duration time.Duration) {
// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("handle", ctx.Err())
// 	case <-time.After(duration):
// 		fmt.Println("process request with", duration)
// 	}
// }

func main() {
	m := make(map[int]any)
	m[1] = []int{1, 2}
	ctx := context.WithValue(context.Background(), 1, m[1])
	handle(ctx)
}

func handle(ctx context.Context) {
	fmt.Println(ctx.Value(1))
}
