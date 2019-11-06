package main

import (
	"testing"
)
func BenchmarkInterfaceYes(b *testing.B) {
	c := NewCaller()
	for i := 0; i < b.N; i++ {
		c.Call()
	}
}
func BenchmarkInterfaceNo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Call()
	}
}
