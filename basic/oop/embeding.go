package main

import "fmt"

func main() {
	parent := Parent{
		Name: "Tree",
		Age:  59,
	}
	parent.SayHello()
	fmt.Println(parent.ToString())

	Parent{}.SayHello()
	//fmt.Println(Parent{}.ToString())


}

type Parent struct {
	Name string
	Age int
}

func (p *Parent) ToString() string{
	return fmt.Sprintf("Name: %s, Age %d", p.Name, p.Age)
}

func(p Parent) SayHello() {
	fmt.Println("Hello I am parent", p.Name)
}

type FirstChild struct {
	Parent
}

type SecondChild struct{
	*Parent
}

type ThirdChild struct{
	Parent
}


func (p *ThirdChild) ToString() string{
	return fmt.Sprintf("Name: %s, Age %d", p.Name, p.Age)
}

func(p ThirdChild) SayHello() {
	fmt.Println("Hello I am child", p.Name)
}