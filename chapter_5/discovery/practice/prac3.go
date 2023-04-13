package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type MyStruct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (m *MyStruct) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID   json.RawMessage `json:"id"`
		Name string          `json:"name"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if num, err := strconv.ParseInt(string(aux.ID), 10, 64); err == nil {
		m.ID = strconv.FormatInt(num, 10)
	} else {
		m.ID = string(aux.ID)
	}

	m.Name = aux.Name
	return nil
}

func main() {
	// JSON 직렬화
	obj := MyStruct{
		ID:   "1234567890123456789",
		Name: "John",
	}
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(b))

	// JSON 역직렬화
	var obj2 MyStruct
	if err := json.Unmarshal(b, &obj2); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(obj2)
}
