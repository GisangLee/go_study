package jay

import "strings"

type CaseInsensitive []string

func (c CaseInsensitive) Len() int {
	return len(c)
}

func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) ||
		(strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])
}

func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// DeadlineAsc 연습문제 4
type DeadlineAsc []Task

func (d DeadlineAsc) Len() int {
	return len(d)
}

func (d DeadlineAsc) Less(i, j int) bool {
	return d[i].Deadline.Time.Before(d[j].Deadline.Time)
}

func (d DeadlineAsc) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// PriorityAsc 연습문제 4
type PriorityAsc []Task

func (p PriorityAsc) Len() int {
	return len(p)
}

func (p PriorityAsc) Less(i, j int) bool {
	return p[i].Priority < p[j].Priority
}

func (p PriorityAsc) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
