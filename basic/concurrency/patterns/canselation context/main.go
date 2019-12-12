package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++{
		go workerWithTimout(ctx, ch)
	}
	cancel()
	time.Sleep(time.Second * 7)
}

func workerWithTimout(ctx context.Context, ch chan int){
loop: for{
	select {
	case <-ch:
		// a read from ch has occurred
	case <-ctx.Done():
		fmt.Println("Stop loop")
		break loop
	}
}
	fmt.Println("Stop worker")
}
