package jay

import (
	"fmt"
	"sort"
)

func ExampleCount() {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	var keys sort.IntSlice // []int (sort 와 관련된 메서드를 가짐)
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)
	for _, key := range keys {
		fmt.Println(string(rune(key)), codeCount[rune(key)])
	}
	// Output:
	// 가 1
	// 나 2
	// 다 1
}
