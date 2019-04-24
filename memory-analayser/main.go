package main

import (
	"fmt"
	"github.com/capnm/sysinfo"
	"github.com/shirou/gopsutil/mem"
	"log"
	"runtime"
)

func main() {
	CheckMemoryOldGitRepo()
	fmt.Println("--------------------------------")
	PrintMemUsage()
	fmt.Println("--------------------------------")
	CheckMemoryGopsutil()
}

// A MemStats records statistics about the memory allocator.
// For info on each, see: https://golang.org/pkg/runtime/#MemStats
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// the number of logical CPUs usable by the current process.
	log.Println("Number of CPUs: ", runtime.NumCPU())
	// the number of goroutines that currently exist.
	log.Println("Count gorutine: ", runtime.NumGoroutine())

	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\nTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\nSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\nNumGC = %v", m.NumGC)
	fmt.Printf("\nHeapSys = %v\n", m.HeapSys)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// There I am using external library "github.com/capnm/sysinfo"
func CheckMemoryOldGitRepo(){
	si := sysinfo.Get()
	fmt.Println("Total RAM :", si.TotalRam)
	fmt.Println("Free RAM :", si.FreeRam)
}

func CheckMemoryGopsutil(){
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v \nFree:%v \nUsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}