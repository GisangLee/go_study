package main

import "fmt"

func sum(a, b int, c chan int) {
	c <- a + b //채널에 a와 b의 합을 보냄.
	//<-c 는 채널에서 값이 들어올때까지 대기. 그리고 채널에 값이 들어오면 대기를 끝내고 다음 코드 실행.
}

func main() {

	//채널은 고루틴끼리 데이터를 주고받고 실행 흐름을 제어하는 기능. 모든 타입을 채널로 사용가능 . 채널은 레퍼런스 타입
	//c는 동기(synchronous) 채널
	c := make(chan int)

	go sum(1, 2, c)

	//채널에 값이 들어올 때까지 기다린 뒤 값이 들어오면 값을 가져옴.
	n := <-c

	//값이 들어오면 다음 코드 실행
	fmt.Println(n)
}
