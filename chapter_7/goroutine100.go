package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	//랜덤 숫자 생성
	r := rand.Intn(30 * 1000 * 1000 * 1000)
	//랜덤 시간 대기
	time.Sleep(time.Duration(r))
	fmt.Println(n)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i)
	}
	//Scanln 으로 main 함수 종료 대기
	fmt.Scanln()
}
