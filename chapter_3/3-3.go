package main

import "fmt"
import "sort"

func search_string(strings []string, searching_word string) {
	for i, word := range strings {
		if word == searching_word {
			fmt.Printf("'%s' found!! (index:[%d])\n", searching_word, i)
			return
		}
	}
	
	fmt.Printf("'%s' not found.\n", searching_word)
}

func main() {
	strings := []string{"hello", "abcd", "world", "korea"}

	sort.Strings(strings)
	fmt.Println(strings)

	searching_word := "korea"
	fmt.Printf("search at word '%s'\n", searching_word)
	search_string(strings, searching_word)
	
	searching_word = "hi"
	fmt.Printf("search at word '%s'\n", searching_word)
	search_string(strings, searching_word)
}
