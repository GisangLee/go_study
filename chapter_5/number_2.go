package chapter5

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

type CaseInsensitive []string

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop(x interface{})
}

// Len 오버라이드
func (c CaseInsensitive) Len() int { return len(c) }

// Less 오버라이드
func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) || strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j]
}

// Swap 오버라이드
func (c CaseInsensitive) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

// Push 오버라이드
func (c *CaseInsensitive) Push(x interface{}) {
	*c = append(*c, x.(string))
}

// Pop 오버라이드
func (c *CaseInsensitive) Pop() interface{} {
	len := c.Len()
	last := (*c)[len-1]
	*c = (*c)[:len-1]
	return last
}

// 기존 정렬
func ExampleCaseInsensitiveSort() {
	apple := CaseInsensitive([]string{"iPone", "iPad", "MacBook", "AppStore"})

	sort.Sort(apple)
}

// 테이블 기반 테스트
func TestCaseInsensitiveSort(t *testing.T) {
	testCase := []struct {
		name     string
		input    CaseInsensitive
		expected CaseInsensitive
	}{
		{
			name:     "빈 슬라이스 정렬",
			input:    CaseInsensitive{},
			expected: CaseInsensitive{},
		},
		{
			name:     "요소 1개만 있는 슬라이스 정렬",
			input:    CaseInsensitive([]string{"apple"}),
			expected: CaseInsensitive([]string{"apple"}),
		},
		{
			name:     "두 요소가 있는 슬라이스 정렬",
			input:    CaseInsensitive([]string{"banana", "Apple"}),
			expected: CaseInsensitive([]string{"Apple", "banana"}),
		},
		{
			name:     "대소문자만 다른 같은 모양새의 요소 정렬",
			input:    CaseInsensitive([]string{"apple", "Apple", "BANANA", "banana"}),
			expected: CaseInsensitive([]string{"apple", "Apple", "BANANA", "banana"}),
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			// 원본 손상을 막기 위한 복사 진행
			inputCopy := make(CaseInsensitive, len(tt.input))
			copy(inputCopy, tt.input)

			sort.Sort(inputCopy)

			// 테스트 케이스에서 예상한 결과 값과 inputCopy의 정렬 결과가 동일한지 검증
			if !reflect.DeepEqual(inputCopy, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, inputCopy)
			}
		})
	}
}
