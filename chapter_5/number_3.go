// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
)

type ID int64

type Data struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}

/*
----------------------------------------------------------------
JSON으로부터 값을 읽어올 때,
먼저 int64 형으로 읽어보고 int64가 아니면
string 형으로 읽는다.
----------------------------------------------------------------
*/
func (id *ID) UnmarshalJSON(data []byte) error {
	var num int64

	if err := json.Unmarshal(data, &num); err == nil {
		*id = ID(num)
		return nil
	}

	var str string

	if err := json.Unmarshal(data, &str); err == nil {
		*id = ID(len(str))
		return nil
	}

	return fmt.Errorf("invalid ID value")
}

/*
----------------------------------------------------------------
MarshalJSON은 string으로 직렬화한다.
----------------------------------------------------------------
*/
func (id ID) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("\"%d\"", id)
	return []byte(str), nil
}

func number_3() {
	// Serialize
	data := Data{
		ID:   ID(1234567890123456789),
		Name: "장고",
	}

	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	// Deserialize
	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(err)
	}

	// %#v -> 해당 값을 생성하는 소스코드 스니펫 출력.
	fmt.Printf("%#v\n", data)
}

func main() {
	fmt.Println("Hello, 世界")
	number_3()
}
