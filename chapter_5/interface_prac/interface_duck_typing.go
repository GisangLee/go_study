package main

import "fmt"

type Duck struct {
}

func (d Duck) quack() {
	fmt.Println("꽥")
}

func (d Duck) feathers() {
	fmt.Println("오리털 따뜻")
}

type Person struct {
}

func (p Person) quack() {
	fmt.Println("사람 꽥")
}

func (p Person) feathers() {
	fmt.Println("사람털 안따듯")
}

type Quacker interface {
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}

func main() {
	var donald Duck
	var victor Person

	inTheForest(donald)
	inTheForest(victor)
}