package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(1)
	//채널에 버퍼를 생성하면 비동기 채널이 됨. 버퍼가 가득차면 값을꺼내서 출력
	done := make(chan int, 2)
	count := 10

	go func() {
		for i := 0; i < count; i++ {
			done <- i
			fmt.Println("고루틴:", i)
		}
	}()

	for i := 0; i < count; i++ {
		s := <-done
		fmt.Println("메인함수:", s)
	}

	/*

	 */
}
