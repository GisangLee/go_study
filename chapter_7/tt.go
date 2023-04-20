package main

import "fmt"

func main() {
	go func() {
		fmt.Println("in goroutine")
	}()
	fmt.Println("in main route")
}
