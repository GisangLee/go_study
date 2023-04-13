package jay

import (
	"testing"
	"time"
)

func TestTask_MarkDone(t *testing.T) {
	subTask1 := Task{
		Title:  "subTask1",
		Status: UNKNOWN,
		Deadline: &Deadline{
			time.Now(),
		},
		Priority: 1,
	}
	subTask2 := Task{
		Title:  "subTask2",
		Status: UNKNOWN,
		Deadline: &Deadline{
			time.Now(),
		},
		Priority: 2,
	}
	mainTask := Task{
		Title:  "mainTask",
		Status: UNKNOWN,
		Deadline: &Deadline{
			time.Now(),
		},
		Priority: 2,
		SubTasks: []Task{subTask1, subTask2},
	}
	mainTask.MarkDone()
	if task, ok := searchStatus(mainTask); !ok {
		t.Errorf("%s status => %v", task.Title, task.Status)
	}
}

func searchStatus(task Task) (*Task, bool) {
	if task.Status != DONE {
		return &task, false
	}
	for _, t := range task.SubTasks {
		if _, ok := searchStatus(t); !ok {
			return &t, false
		}
	}
	return nil, true
}
