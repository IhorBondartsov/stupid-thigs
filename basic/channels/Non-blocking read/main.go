package main

import (
	"fmt"
	"time"
)

func main() {
	myChan := make(chan string)

	go func() {
		myChan <- "Message!"
	}()

	// go Read(myChan) // не самый лучший вариант так как идет бесконечное итерирование по селекту
	go ReadWithRange(myChan)

	time.Sleep(2 * time.Second)

	go func() {
		myChan <- "Message!"
	}()
	time.Sleep(time.Second * 2)
	close(myChan)
	time.Sleep(time.Second * 2)
}

// bad variant all time will be printed "No Msg"
func Read(myChan chan string) {
End:
	for {
		select {
		case msg, ok := <-myChan:
			if !ok {
				break End
			}
			fmt.Println(msg)
		default:
			fmt.Println("No Msg")
		}
	}
	fmt.Println("Stop")
}

func ReadWithRange(myChan chan string) {
	for v := range myChan {
		fmt.Println(v)
	}
	fmt.Println("Stop")
}
