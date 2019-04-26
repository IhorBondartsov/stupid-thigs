package main

import "fmt"

// Вот так выглядит слайс внутри
//			type slice struct {
//				array unsafe.Pointer
//				len   int
//				cap   int
//			}
// Масивы передаются по значению - слайсы по указателю
// При создании слайса make([]int, 5, 10)  мы получаепм в слайсе пять нулей
// и пять зарезервированых мест для вставки элементов

func main() {
	// Create slice
	// len 5 cap 10
	// 5 elements will be 0 and other not init
	slice := make([]int, 5, 10)
	fmt.Println(slice[0]) // 0
	// fmt.Println(slice[8]) // panic because element number 8 not init

	// ----------Second part: Slice and array
	arr := [3]int{1,2,3}
	slice2 := []int{1,2,3}

	ChangeArray(arr)
	ChangeSlice(slice2)

	fmt.Println("Array",arr) // 1,2,3
	fmt.Println("Slice", slice2) // 1,10,3

	fmt.Println(arr == arr) // true
	//fmt.Println(slice2 == slice2) // Doesnt work compile error
}

func ChangeSlice(slice []int){
	slice[1] = 10
}

func ChangeArray(arr [3]int){
	// Doesnt work next
	// func ChangeArrayError(arr [...]int)
	// func ChangeArrayError(arr [50]int)
	// func ChangeArrayError(arr []int)
	arr[1] = 10
}



