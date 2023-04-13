package jay

import (
	"errors"
	"strconv"
	"time"
)

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

type status int

// MarshalJSON implements the json.Marshaler interface
func (s status) MarshalJSON() ([]byte, error) {
	switch s {
	case UNKNOWN:
		return []byte(`"UNKNOWN"`), nil
	case TODO:
		return []byte(`"TODO"`), nil
	case DONE:
		return []byte(`"DONE"`), nil
	default:
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *status) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"UNKNOWN"`:
		*s = UNKNOWN
	case `"TODO"`:
		*s = TODO
	case `"DONE"`:
		*s = DONE
	default:
		return errors.New("status.UnmarshalJSON: unknown value")
	}
	return nil
}

type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

func (d Deadline) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, d.Unix(), 10), nil
}

func (d *Deadline) UnmarshalJSON(data []byte) error {
	unix, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.Unix(unix, 0)
	return nil
}

// MarkDone 연습문제 1
func (t *Task) MarkDone() {
	if t.Status == DONE {
		return
	}

	t.Status = DONE

	if len(t.SubTasks) == 0 {
		return
	}
	for i := range t.SubTasks {
		t.SubTasks[i].Status = DONE
	}
}
