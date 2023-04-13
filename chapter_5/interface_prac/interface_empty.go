package main

import (
	"fmt"
	"strconv"
)

type Empty interface {
}

type Person struct{
	age int
	name string
}

func EmptyFunc(arg interface{}) {
	fmt.Println("hi")
}

func TypeChecker(arg interface{}) string {
	//Type assertion
	// switch 분기분 안에서만 .(type) 사용가능
	switch arg.(type){
	case int:
		i:=arg.(int)
		return strconv.Itoa(i)
	case float32:
		i:=arg.(float32)
		return strconv.FormatFloat(float64(i),'f',-1,32)
	
	//구조체 값이 넘어올경우
	case Person:
		i:=arg.(Person)
		return i.name +" " +strconv.Itoa(i.age)	
	default:
		return "error"
	}
}


func main(){

	empty_test:=3
	EmptyFunc(empty_test)

	fmt.Println(TypeChecker(1))
	
	fmt.Println(TypeChecker(1.5))
	
	fmt.Println(TypeChecker("안녕하세유"))

	person:=Person{name: "김개똥",age: 50}
	fmt.Println(TypeChecker(person))

}