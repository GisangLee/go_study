package jay

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"time"
)

func ExampleCaseInsensitive_sort() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	sort.Sort(apple)
	fmt.Println(apple)
	// Output:
	// [AppStore iPad iPhone MacBook]
}

// TestCaseInsensitiveSort 연습문제 2
func TestCaseInsensitiveSort(t *testing.T) {
	tests := []struct {
		name string
		in   CaseInsensitive
		want []string
	}{
		{
			"test1",
			CaseInsensitive([]string{
				"iPhone", "iPad", "MacBook", "AppStore",
			}),
			[]string{"AppStore", "iPad", "iPhone", "MacBook"},
		},
		{
			"test2",
			CaseInsensitive([]string{
				"iPhone", "iPad", "IPad", "MacBook", "AppStore",
			}),
			[]string{"AppStore", "IPad", "iPad", "iPhone", "MacBook"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(tt.in)
			for i := range tt.in {
				if tt.in[i] != tt.want[i] {
					t.Errorf("Sort() = %+v, want %+v", tt.in, tt.want)
					break
				}
			}
		})
	}
}

func TestDeadlineAscSort(t *testing.T) {
	testTime1 := time.Now()
	testTime2 := time.Now()
	testTime3 := time.Now()
	tests := []struct {
		name string
		in   DeadlineAsc
		want []Task
	}{
		{
			"task2-task1-task3",
			DeadlineAsc([]Task{
				{
					Title:    "task1",
					Deadline: NewDeadline(testTime2),
				},
				{
					Title:    "task2",
					Deadline: NewDeadline(testTime1),
				},
				{
					Title:    "task3",
					Deadline: NewDeadline(testTime3),
				},
			}),
			[]Task{
				{
					Title:    "task2",
					Deadline: NewDeadline(testTime1),
				},
				{
					Title:    "task1",
					Deadline: NewDeadline(testTime2),
				},
				{
					Title:    "task3",
					Deadline: NewDeadline(testTime3),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(tt.in)
			for i := range tt.in {
				if !reflect.DeepEqual(tt.in[i], tt.want[i]) {
					t.Errorf("Sort() = %+v, want %+v", tt.in, tt.want)
					break
				}
			}
		})
	}
}

func TestPriorityAscSort(t *testing.T) {
	tests := []struct {
		name string
		in   PriorityAsc
		want []Task
	}{
		{
			"task2-task1-task3",
			PriorityAsc([]Task{
				{
					Title:    "task1",
					Priority: 2,
				},
				{
					Title:    "task2",
					Priority: 1,
				},
				{
					Title:    "task3",
					Priority: 3,
				},
			}),
			[]Task{
				{
					Title:    "task2",
					Priority: 1,
				},
				{
					Title:    "task1",
					Priority: 2,
				},
				{
					Title:    "task3",
					Priority: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(tt.in)
			for i := range tt.in {
				if !reflect.DeepEqual(tt.in[i], tt.want[i]) {
					t.Errorf("Sort() = %+v, want %+v", tt.in, tt.want)
					break
				}
			}
		})
	}
}
