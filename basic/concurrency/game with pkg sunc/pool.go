package main

import (
	"fmt"
	"sync"
)

// Pool for our struct A
var pool *sync.Pool

// A dummy struct with a member
type A struct {
	Name string
}

// Func to init pool
func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("Returning new A")
			return new(A)
		},
	}
}

// Main func
func main() {
	// Initializing pool
	initPool()
	// Get hold of instance one
	one := pool.Get().(*A)
	one.Name = "first"
	fmt.Printf("one.Name = %s\n", one.Name)
	// Submit back the instance after using
	pool.Put(one)
	// Now the same instance becomes usable by another routine without allocating it again

	two := pool.Get().(*A)
	fmt.Printf("one.Name = %s\n", two.Name)
}
