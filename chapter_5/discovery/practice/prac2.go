package main

import (
	"fmt"
	"sort"
	"strings"
)

type CaseInsensitive []string

// 배열 크기
func (c CaseInsensitive) Len() int {
	return len(c)
}

func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) || (strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])
}

// Swap swaps the elements with indexes i and j.
func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *CaseInsensitive) Push(x interface{}) {
	*c = append(*c, x.(string))
}

func (c *CaseInsensitive) Pop() interface{} {
	len := c.Len()
	last := (*c)[len-1]
	*c = (*c)[:len-1]
	return last
}

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

// 연습문제 2
// 테이블 기반 테스트
func TestCaseInsensitiveSort() {
	cases := []CaseInsensitive{{"iPhone", "iPad", "MacBook", "AppStore"}, {"hi", "bye"}, {"c", "D", "A", "j"}}

	for i, _ := range cases {
		sort.Sort(cases[i])
		fmt.Println(cases[i])
	}
}

func main() {

	TestCaseInsensitiveSort()
	//
	//heap.Init(&apple)
	//for apple.Len() > 0 {
	//	fmt.Println(heap.Pop(&apple))
	//}
	//sort.Sort(apple)
	//fmt.Println(apple)
}
