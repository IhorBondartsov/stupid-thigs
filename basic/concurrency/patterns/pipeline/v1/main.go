package main

import "fmt"

func main() {
	res := <-pipe(Add(1,2,3,4,5), Minus, 2,3,4,5)
	if res != 1 {
		fmt.Println("Something wrong", res)
	}
}

func pipe(left <-chan int, f func(i ...int) <-chan int, i ...int) <-chan int {
	is := []int{<-left}
	is = append(is, i...)
	return f(is...)
}

func Add(i ...int) <-chan int {
	out := make(chan int)
	go func() {
		res := 0
		for _, v := range i {
			res += v
		}
		out <- res
	}()
	return out
}

func Minus(i ...int) <-chan int {
	out := make(chan int)
	go func() {
		res := 0
		for k, v := range i {
			if k == 0{
				res += v
				continue
			}
			res -= v
		}
		fmt.Println(res)
		out <- res
	}()
	return out
}
