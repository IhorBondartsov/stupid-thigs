package main

import (
	"fmt"
	"sync"
)

//sync.Pool [1] comes really handy when one wants to reduce the number of allocations happening during the course of a
// functionality written in Golang. A Pool is a set of temporary objects that may be individually saved and retrieved.
// fasthttp [2], zerolog [3] are couple of those most popuplar open source Golang libraries which uses sync.Pool
// at the core of their implementation.


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
