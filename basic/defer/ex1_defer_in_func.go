package main

import "fmt"

// Демонстрирует что все деверы находятся в LIFO
func startEx1(){
	fmt.Println("Start startEx1")
	defer Func1()
	Func2()
	fmt.Println("End startEx1")
}

func Func1(){
	defer func(){
		fmt.Println("defer in func1")
	}()
}

func Func2(){
	fmt.Println("Start func2")
	defer func(){
		fmt.Println("defer in func2")
	}()
	fmt.Println("End func2")
}

func CalledFuncWithDefer(){
	Func2()
}
