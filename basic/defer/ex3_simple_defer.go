package main

import "fmt"

// print:
// 2
// 1
// 0

func startEx3(){
	for i:=0; i < 3; i++{
		defer fmt.Println(i)
	}
}