package jay

import (
	"errors"
	"strconv"
	"strings"
)

type Exercise3 struct {
	Id ID `json:"id"`
}

type ID int64

type Test interface {
}

func (id *ID) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return errors.New("id.UnmarshalJSON: unknown value")
	}
	*id = ID(i)
	return nil
}
