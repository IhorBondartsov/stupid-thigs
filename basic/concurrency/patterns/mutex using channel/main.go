package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := NewMap()

	for _, v := range arr {
		go m.Add(v)
	}

	time.Sleep(time.Second * 5)
	for _, v := range arr[1:] {
		go m.Delete(v)
	}
	time.Sleep(time.Second * 5)
	m.print()
}

type myMap struct {
	m     map[int]struct{}
	mutex chan struct{}
}

func (m *myMap) Lock() {
	<-m.mutex
}

func (m *myMap) Unlock() {
	m.mutex <- struct{}{}
}

func NewMap() *myMap {
	m := &myMap{
		m:     map[int]struct{}{},
		mutex: make(chan struct{}, 1),
	}
	m.Unlock()
	return m
}

func (m *myMap) Add(i int) {
	m.Lock()
	defer m.Unlock()
	m.m[i] = struct{}{}

}

func (m *myMap) Delete(i int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, i)
}

func (m *myMap) print() {
	m.Lock()
	defer m.Unlock()
	fmt.Println(m.m)
}
