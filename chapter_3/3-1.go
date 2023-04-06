package main

import "fmt"

func add_suffix(word string) string {
	var word_arr = []rune(word)

	if (word_arr[len(word_arr) - 1] - 44032) % 28 == 0 {
		word += "는"
	} else {
		word += "은"
	}
		
	return word
}

func main() {
    fruits := []string{"사과", "바나나", "토마토", "멜론"}
    
	for _, fruit := range fruits {
		fmt.Printf("%s 맛있다.\n", add_suffix(fruit))
	}
}
