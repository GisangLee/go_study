package jay

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func Example_marshalJSON() {
	t := Task{
		Title:    "Laundry",
		Status:   DONE,
		Deadline: NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output:
	// {"title":"Laundry","status":"DONE","deadline":1439739780}
}

func Example_unmarshalJSON() {
	b := []byte(`
	{
		"title":"Buy Milk",
		"status":"DONE",
		"deadline":1439739780
	}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
	// Output:
	// Buy Milk
	// 2
	// 2015-08-16 15:43:00 +0000 UTC
}

func Example_marshalJSON_Exercise3() {
	e := Exercise3{Id: 1}
	b, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output:
	// {"id":"1"}
}

func Example_unmarshalJSON_Exercise3() {
	b1 := []byte(`
	{
		"id":"1"
	}`)
	b2 := []byte(`
	{
		"id":2
	}`)
	e1 := Exercise3{}
	e2 := Exercise3{}
	err := json.Unmarshal(b1, &e1)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(b2, &e2)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(e1.Id)
	fmt.Println(e2.Id)
	// Output:
	// 1
	// 2
}
