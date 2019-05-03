package main

import "fmt"

func startEx2(){
	var i int
	defer fmt.Println(i) // print 0
	defer func(){
		fmt.Println(i) // print 10
	}()
	i = 10
}
