// go run condition.go
//
// A condition variable is basically a container of threads that are waiting for a certain condition.
//
// Go provides condition variable via sync.Cond
// Condition variables are useful if all you want is to signal other goroutines that some event has occured.
// Use Signal() to notify one of the goroutines that is waiting for the longest period(Go uses FIFO data structure internally to keep track waiting goroutines on a condition variable).
// Condition variables are perfect fit if you need to broadcast that some event has occured to all the goroutines waiting for that particular event.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type Record struct {
	sync.Mutex

	buf  string
	cond *sync.Cond

	writers []io.Writer
}

func NewRecord(writers ...io.Writer) *Record {
	r := &Record{writers: writers}
	r.cond = sync.NewCond(r)
	return r
}

func (r *Record) Prompt() {
	for {
		var s string

		fmt.Printf(":> ")
		fmt.Scanf("%s", &s)

		if s == "exit" {
			return
		}

		r.Lock()
		r.buf = s
		r.Unlock()

		r.cond.Broadcast()
	}
}

func (r *Record) Start() error {
	f := func(w io.Writer) {
		for {
			r.Lock()
			r.cond.Wait()
			fmt.Fprintf(w, "%s\n", r.buf)
			r.Unlock()
		}
	}
	for i := range r.writers {
		go f(r.writers[i])
	}
	return nil
}

func main() {
	f, err := os.Create("cond.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f2, err := os.Create("cond2.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	r := NewRecord(f, f2)
	r.Start()
	r.Prompt()
}
