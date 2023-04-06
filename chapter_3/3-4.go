package main

import "fmt"

func enqueue(q []string, str string) []string {
	return append(q, str)
}

func dequeue(q []string) []string {
	return q[1:]
}

func main() {
	q := []string{}

	word1 := "good"
	word2 := "day"
	word3 := "commander"
	
	fmt.Printf("enqueue '%s'\n", word1)
	q = enqueue(q, word1)
	fmt.Println(q)

	fmt.Printf("enqueue '%s'\n", word2)
	q = enqueue(q, word2)
	fmt.Println(q)

	fmt.Printf("enqueue '%s'\n", word3)
	q = enqueue(q, word3)
	fmt.Println(q)
	
	q = dequeue(q)
	fmt.Println("dequeue")
	fmt.Println(q)

	q = dequeue(q)
	fmt.Println("dequeue")
	fmt.Println(q)

	q = dequeue(q)
	fmt.Println("dequeue")
	fmt.Println(q)

	// q = dequeue(q) // out of range 로 인한 패닉 발생함. size 선체크후 처리가 필요할 것으로 판단됨.
}
