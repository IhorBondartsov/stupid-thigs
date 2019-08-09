// В этом примере такая ситуация:
// Герои общаються по каналам а шпион слушает все и передат информацию по каналу Саурону
package main

import (
	"fmt"
	"sync"
	"time"
)

type human struct{
	voice chan string
	name string
}

func (h *human) SayName(){
	h.voice <- h.name
}

func NewHuman(name string)*human{
	return &human{
		voice: make(chan string),
		name:name,
	}
}

// FUN-IN
func MordorsSpyListenAndSayToEye( voices ...chan string)<-chan string{
	var wg sync.WaitGroup
	out := make(chan string)
	for _, value := range voices {
		v := value
		wg.Add(1)
		go func(){
			for i := range v {out <- i}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	crowd := []*human{NewHuman("Gendalf"), NewHuman("Aragorn"), NewHuman("Frodo")}

	// spy found all heroes
	heroes := []chan string{}
	for _, value := range crowd {
		heroes = append(heroes, value.voice)
	}
	// spy is starting listen all heroes
	sauronListen := MordorsSpyListenAndSayToEye(heroes...)

	// heroes are speaking
	for _, value := range crowd {
		value := value
		go func(){
			time.Sleep(time.Second * 1)
			value.SayName()
			close(value.voice)
		}()
	}

	// Sauron know everything
	for i:= 0; i < 3; i++{
		fmt.Println(<-sauronListen)
	}

}