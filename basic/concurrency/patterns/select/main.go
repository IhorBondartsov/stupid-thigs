package main

import (
	"fmt"
	"time"
)
// В этом примере просто демонстрируеться как передаеться информация в канал
func main() {
	ch := make(chan int)
	done := make(chan struct{})

	go  Print(ch, done)

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			close(done)
			return
		}
	}

	time.Sleep(time.Second * 5)
}

func Print(to chan int, done chan struct{}) {
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			select {
			// вот тут вся красота
			case to <- i:
			case <-done:
				fmt.Println("done", i)
				return
			}
		}()
	}
}
