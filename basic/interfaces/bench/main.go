package main

type Caller interface {
	Call()
}

func NewCaller()Caller{
	return caller{}
}

type caller struct {}

func (c caller) Call() {}

func Call(){}

