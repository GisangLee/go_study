package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type status int
type Deadline time.Time

type Task struct{
	Title string
	Status status
	deadline *Deadline

}

func (d *Deadline) OverDue() bool{
	return d !=nil && time.Time(*d).Before(time.Now())
}


func (t Task) OverDue() bool {

	return t.deadline.OverDue()
}

func main() {
	d1 := Deadline(time.Now().Add(-4 * time.Hour))
	d2 := Deadline(time.Now().Add(4 * time.Hour))
	d3 := Deadline(time.Now().Add(-4 * time.Hour))

	t1 := Task{"4h ago",1,&d1}
	t2 := Task{"4h later",1,&d2}
	t3 := Task{"4h later",1,&d3}
	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())

	//marshaling
	
	b,err:=	json.Marshal(t3)

	if err != nil{
		fmt.Println(err)
		return 
	}

	fmt.Println(string(b))

}