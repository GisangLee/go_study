package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool) //동기 채널
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			//고루틴에 true를 보냄, 값을 꺼낼 때 까지 대기
			done <- true
			fmt.Println("고루틴:", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		//채널에 값이 들어올때까지 대기, 값을 꺼냄
		a := <-done
		fmt.Println("메인함수:", i, a)

	}

	/*
		결과값
			고루틴: 0
			메인함수: 0 true
			고루틴: 1
			메인함수: 1 true
			고루틴: 2
			메인함수: 2 true

	*/
}
