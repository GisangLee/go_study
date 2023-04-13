package main

import "fmt"

type MyInt int

func (i MyInt) Print(){
	fmt.Println(i)
}

type Printer interface{
	Print()
}

func main(){
	var i MyInt =5
	var p Printer

p=i

p.Print()
}