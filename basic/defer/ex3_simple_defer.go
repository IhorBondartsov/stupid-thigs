package main

import "fmt"

// print:
// 2
// 1
// 0

func startEx3(){
	s1 := []int{1, 2, 3, 4} // 4 4
	s2 := s1[0:3] // 1 2 3  // 3 4

	s2 = append(s2, 6) // 1 2 3 6  // 4, 4

	s1[2] = 10 // 1 2 10 4

	fmt.Println(s1)
	fmt.Println(s2)
}