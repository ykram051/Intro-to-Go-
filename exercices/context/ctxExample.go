package main

import (
	"context"
	"fmt"
	"time"
)

func dowork(ctx context.Context) {
	time.Sleep(time.Second)
	fmt.Println("hello")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
	}
}

func main() {
	ctx ,cancel:= context.WithCancel(context.Background())

	go dowork(ctx)
	time.Sleep(5* time.Second)
	cancel()
	time.Sleep(5* time.Second)

	// Create a context with a timeout of 2 seconds
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Simulate a long-running task
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task cancelled:", ctx.Err())
	}
}