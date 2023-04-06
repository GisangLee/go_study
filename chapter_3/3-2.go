package main

import "fmt"
import "sort"

func slice_sort(ints []int) []int {
	sort.Ints(ints)

	return ints
}

func main() {
	ints := []int{4,8,5,1,2,6,9,3,7}
	
	fmt.Println(ints)
	slice_sort(ints)
	fmt.Println(ints)
}
