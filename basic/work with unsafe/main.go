package main

import (
	"fmt"
	"unsafe"
)

// unsafe - не гарантирует обратную совместимость

func main() {
	// Смотрим размер и указатель переменной

	var x int64 = 955835
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Pointer(&x))

	// Смотрим кишки структуры
	type Object struct {
		Valid           bool
		Value           int
		Name            string
		MetaInformation map[string]string
		Childs          []*Object
	}

	obj := Object{
		Name: "Big dady",
		Value:5,
		MetaInformation: map[string]string{
			"age":"50",
			"time": "19:40",
		},
		Valid:true,
		Childs: []*Object{
			&Object{
				Name:"sister",
			},
		},
	}

	fmt.Println("Object size: ", unsafe.Sizeof(obj))

	fmt.Println(
		"Name: SizeOf", unsafe.Sizeof(obj.Name),
		"Alignof", unsafe.Alignof(obj.Name),
		"OffsetOff", unsafe.Offsetof(obj.Name),
		)

	fmt.Println(
		"Valid: SizeOf", unsafe.Sizeof(obj.Valid),
		"Alignof", unsafe.Alignof(obj.Valid),
		"OffsetOff", unsafe.Offsetof(obj.Valid),
	)

	fmt.Println(
		"Value: SizeOf", unsafe.Sizeof(obj.Value),
		"Alignof", unsafe.Alignof(obj.Value),
		"OffsetOff", unsafe.Offsetof(obj.Value),
	)

	fmt.Println(
		"MetaInformation: SizeOf", unsafe.Sizeof(obj.MetaInformation),
		"Alignof", unsafe.Alignof(obj.MetaInformation),
		"OffsetOff", unsafe.Offsetof(obj.MetaInformation),
	)

	fmt.Println(
		"Childs: SizeOf", unsafe.Sizeof(obj.Childs),
		"Alignof", unsafe.Alignof(obj.Childs),
		"OffsetOff", unsafe.Offsetof(obj.Childs),
	)
}
