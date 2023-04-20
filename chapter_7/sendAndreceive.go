package main

import "fmt"

// 보내기 전용, 받기 전용 채널 테스트
func producer(c chan<- int) { //보내기 전용 채널
	for i := 0; i < 5; i++ {
		c <- i
	}
	c <- 100
	//fmt.Println(<-c) 채널에서 값을 꺼내면 컴파일 에러
}

func customer(c <-chan int) { //받기 전용 채널

	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(<-c) //채널 값 꺼내기
	//c<- 1  채널에 값 보내면 컴파일 에러
}
func main() {

	c := make(chan int)

	go producer(c)
	go customer(c)
	fmt.Scanln()
}
