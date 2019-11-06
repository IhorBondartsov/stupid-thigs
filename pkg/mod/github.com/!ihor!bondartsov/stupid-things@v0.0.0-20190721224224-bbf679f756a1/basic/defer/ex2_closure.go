package main

import "fmt"

// переменная i замыкается в превом дефере но во втором нет.
// 10
// 0
func startEx2(){
	var i int
	defer fmt.Println(i) // print 0
	defer func(){
		fmt.Println(i) // print 10
	}()
	i = 10
}
