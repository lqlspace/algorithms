package main

import (
	"fmt"
	"time"
	"context"
)

func f3(ctx context.Context) int {
	fmt.Println(ctx.Value("NLJB"))

	select {
	case <-ctx.Done():
		return 3
	default:

	}
	return 0
}

func f2(ctx context.Context) int {
	fmt.Println(ctx.Value("HELLO"))
	fmt.Println(ctx.Value("WROLD"))

	ctx = context.WithValue(ctx, "NLJB", "NULIJIABEI")

	go fmt.Println(f3(ctx))

	select {
	case <- ctx.Done():
		return 2
	default:
		
	}
	return 0
}

func f1(ctx context.Context) int {
	ctx = context.WithValue(ctx, "HELLO", "WORLD")
	ctx = context.WithValue(ctx, "WORLD", "HELLO")

	go fmt.Println(f2(ctx))

	select {
	case <-ctx.Done():
		return 1
	default:
	}

	return 0
}


func main() {
	timeout := 3 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	fmt.Println(f1(ctx))

/*	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()
	fmt.Println(f1(ctx))*/
}
