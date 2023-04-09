package chapter5

import (
	"time"
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

const (
	UNKNOWN status = 0
	TODO    status = 1
	DONE    status = 2
)

func NewDeadLine(t time.Time) *DeadLine {
	return &DeadLine{t}
}

// 연습문제 1번의 함수
func (t *Task) MakeDone() {
	for idx := range t.SubTasks {
		t.SubTasks[idx].Status = TODO
	}

	/* 중요 !!
	--------------------------------------------------------
	아래 반복문은 SubTasks들의 Task 인스턴스의 값을 바꾸는 것이지
	MkaeDone 메서드가 호출된 t1 인스턴스의 Subtask와는 관련이 없기 떄문에
	값이 변경 되지 않는다.
	--------------------------------------------------------
	*/

	// for _, subtask := range t.SubTasks {
	// 	subtask.Status = DONE
	// }
}

func main() {
	t1 := Task{
		Title:    "헬창",
		Status:   UNKNOWN,
		DeadLine: nil,
		Priority: 1,
		SubTasks: []Task{
			{
				Title:    "서브헬창1",
				Status:   DONE,
				DeadLine: nil,
				Priority: 2,
			},
			{
				Title:    "서브헬창2",
				Status:   TODO,
				DeadLine: nil,
				Priority: 2,
			},
		},
	}

	t1.MakeDone()
}
