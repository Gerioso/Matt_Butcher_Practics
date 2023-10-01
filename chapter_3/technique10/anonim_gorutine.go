package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Foo1")
	go func() {
		fmt.Println("Foo2")
	}()
	fmt.Println("Again Foo 1")
	runtime.Gosched()
}
