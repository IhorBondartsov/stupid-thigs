package main
import (
	"testing"
)
func BenchmarkDeferYes(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doDefer(&t)
	}
}
func BenchmarkDeferNo(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doNoDefer(&t)
	}
}