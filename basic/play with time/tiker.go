package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start time:", time.Now().String())
	tick := time.NewTicker(time.Second * 5)

	i := 0
	loop: for {
		select {
		case t :=  <- tick.C :
			if i == 5 {
				break loop
			}
			fmt.Println(t.String())
			i++
		}
	}
	fmt.Println("End time:", time.Now().String())
}
