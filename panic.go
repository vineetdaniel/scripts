package main

import "fmt"

func main() {
	//panic("PNIC")
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
	
}
