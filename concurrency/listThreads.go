package main

import (
	"fmt"
	"runtime"
)


func noOfThreads() int {
	threads := runtime.GOMAXPROCS(0)
	return threads
}
func main() {
	runtime.GOMAXPROCS(0)
	fmt.Printf("%d  threads available", noOfThreads())
}
