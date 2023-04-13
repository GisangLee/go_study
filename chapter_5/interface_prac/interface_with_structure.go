package main

import "fmt"

type MyInt int

func (i MyInt) Print(){
	fmt.Println(i)
}

type Rectangle struct{
	width,height int
}

func (r Rectangle) Print(){
	fmt.Println(r.height,r.width)
}

type Printer interface{
	Print()
}

func main(){
	var i MyInt =5
	r :=Rectangle{10,20}

	var p Printer
	p=i
	p.Print()

	p=r
	p.Print()

	p1:=Printer(i) //위와 동일. 하지만 초기값을 넣을 수 있음
	p2:=Printer(r) 

	p1.Print()
	p2.Print()

	var j MyInt=10
	k:=Rectangle{10,20}

	pArr := []Printer{j,k} //슬라이스 형태로 저장

	for index,_:=range pArr{
		pArr[index].Print()
	}
	for _,value:=range pArr{
		value.Print()
	}

}