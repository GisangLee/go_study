# 연습 문제 1

```go

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

```


# 연습 문제 2
```go

// 테이블 기반 테스트
func TestCaseInsensitiveSort(t *testing.T) {
	testCases := []struct {
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

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// 원본 손상을 막기 위한 복사 진행
			inputCopy := make(CaseInsensitive, len(testCase.input))
			copy(inputCopy, testCase.input)

			sort.Sort(inputCopy)

			// 테스트 케이스에서 예상한 결과 값과 inputCopy의 정렬 결과가 동일한지 검증
			if !reflect.DeepEqual(inputCopy, testCase.expected) {
				t.Errorf("expected %v, but got %v", testCase.expected, inputCopy)
			}
		})
	}
}

```


# 연습 문제 3
```go
/*
----------------------------------------------------------------
JSON으로부터 값을 읽어올 때,
먼저 int64 형으로 읽어보고 int64가 아니면
string 형으로 읽는다.
----------------------------------------------------------------
*/
func (id *ID) UnmarshalJSON(data []byte) error {
	var num int64

	if err := json.Unmarshal(data, &num); err == nil {
		*id = ID(num)
		return nil
	}

	var str string

	if err := json.Unmarshal(data, &str); err == nil {
		return err
	}

	fmt.Println(str)
	fmt.Println(len(str))

	*id = ID(len(str))
	return nil
}

/*
----------------------------------------------------------------
MarshalJSON은 string으로 직렬화한다.
----------------------------------------------------------------
*/
func (id ID) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("\"%d\"", id)
	return []byte(str), nil
}
```


# 연습 문제 4
```go

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

```