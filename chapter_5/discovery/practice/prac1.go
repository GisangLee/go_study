package main

import (
	"fmt"
	"time"
)

const (
	UNKNOWN status = 0
	TODO    status = 1
	DONE    status = 2
)

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	DeadLine *DeadLine `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subtasks,omitempty"`
}

type status int

type DeadLine struct {
	time.Time
}

func NewDeadLine(t time.Time) *DeadLine {
	return &DeadLine{t}
}

// 연습문제 1
func (t *Task) MarkDone() {
	t.Status = DONE
	if t.SubTasks != nil {
		for i := range t.SubTasks {
			t.SubTasks[i].MarkDone()
		}
	}

}

func main() {
	t1 := Task{
		Title:  "테스트1",
		Status: UNKNOWN,
		SubTasks: []Task{
			{
				Title:  "테스트2",
				Status: UNKNOWN,
				SubTasks: []Task{
					{
						Title:  "테스트4",
						Status: UNKNOWN,
					},
					{
						Title:  "테스트5",
						Status: UNKNOWN,
					},
				},
			},
			{
				Title:  "테스트3",
				Status: UNKNOWN,
				SubTasks: []Task{
					{
						Title:  "테스트6",
						Status: UNKNOWN,
					},
				},
			},
		},
	}
	t1.MarkDone()
	fmt.Println(t1)
}
