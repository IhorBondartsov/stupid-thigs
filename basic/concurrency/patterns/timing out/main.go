//Демонстрация того как можно закрывать каналы по таймауту
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++{
		go workerWithTimout(ch)
	}
	time.Sleep(time.Second * 7)
}

func workerWithTimout(ch chan int){
	loop: for{
		select {
		case <-ch:
			// a read from ch has occurred
		case <-time.After(time.Second * 2):
			fmt.Println("Stop loop")
			break loop
		}
	}
	fmt.Println("Stop worker")
}
