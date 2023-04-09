package chapter5

import (
	"sort"
	"testing"
	"time"
)

type status string

type DeadLine struct {
	time.Time
}

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	DeadLine *DeadLine `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subtasks,omitempty"`
}

// 데드라인 정렬
type DeadlineSorter []Task

// 우선 순위 정렬
type PrioritySorter []Task

const (
	PENDING status = "PENDING"
	DONE    status = "DONE"
	DENY    status = "DENY"
)

/*
----------------------------------------------------------------
데드라인 sort.Interface 구현
----------------------------------------------------------------
*/

func (d DeadlineSorter) Len() int { return len(d) }

func (d DeadlineSorter) Less(i, j int) bool {
	if d[i].DeadLine == nil && d[j].DeadLine == nil {
		return false
	} else if d[i].DeadLine == nil {
		return false
	} else if d[j].DeadLine == nil {
		return true
	}

	return d[i].DeadLine.Before(d[j].DeadLine.Time)
}

func (d DeadlineSorter) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

/*
----------------------------------------------------------------
우선순위 sort.Interface 구현
----------------------------------------------------------------
*/
func (p PrioritySorter) Len() int { return len(p) }

func (p PrioritySorter) Less(i, j int) bool {
	return p[i].Priority > p[j].Priority
}

func (p PrioritySorter) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

/*
----------------------------------------------------------------
데드라인 테스트 케이스
----------------------------------------------------------------
*/
func TestSortByDeadline(t *testing.T) {
	task1 := Task{
		Title:    "Task 1",
		Status:   PENDING,
		DeadLine: &DeadLine{time.Date(2023, time.April, 30, 23, 59, 59, 0, time.UTC)},
		Priority: 2,
	}
	task2 := Task{
		Title:    "Task 2",
		Status:   DONE,
		DeadLine: &DeadLine{time.Date(2023, time.March, 15, 23, 59, 59, 0, time.UTC)},
		Priority: 3,
	}
	task3 := Task{
		Title:    "Task 3",
		Status:   DENY,
		DeadLine: &DeadLine{time.Date(2023, time.February, 1, 23, 59, 59, 0, time.UTC)},
		Priority: 1,
	}
	tasks := []Task{task1, task2, task3}

	expectedTasks := []Task{task3, task2, task1}

	sort.Sort(DeadlineSorter(tasks))
	for i := range tasks {
		if tasks[i].DeadLine.Time != expectedTasks[i].DeadLine.Time {
			t.Errorf("Expected task %d with deadline %s, but got %s", i, expectedTasks[i].DeadLine.Time, tasks[i].DeadLine.Time)
		}
	}
}

/*
----------------------------------------------------------------
우선순위 테스트 케이스
----------------------------------------------------------------
*/
func TestSortByPriority(t *testing.T) {
	task1 := Task{
		Title:    "Task 1",
		Status:   PENDING,
		DeadLine: &DeadLine{time.Date(2023, time.April, 30, 23, 59, 59, 0, time.UTC)},
		Priority: 2,
	}
	task2 := Task{
		Title:    "Task 2",
		Status:   DENY,
		DeadLine: &DeadLine{time.Date(2023, time.March, 15, 23, 59, 59, 0, time.UTC)},
		Priority: 3,
	}
	task3 := Task{
		Title:    "Task 3",
		Status:   DONE,
		DeadLine: &DeadLine{time.Date(2023, time.February, 1, 23, 59, 59, 0, time.UTC)},
		Priority: 1,
	}
	tasks := []Task{task1, task2, task3}

	expectedTasks := []Task{task3, task1, task2}

	sort.Sort(PrioritySorter(tasks))
	for i := range tasks {
		if tasks[i].Priority != expectedTasks[i].Priority {
			t.Errorf("Expected task %d with priority %d, but got %d", i, expectedTasks[i].Priority, tasks[i].Priority)
		}
	}
}
