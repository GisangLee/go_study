package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)    //int 채널 생성
	c2 := make(chan string) //string 채널 생성

	go func() {
		for {
			c1 <- 10                           //채널에 10 보내기
			time.Sleep(100 * time.Millisecond) //100 밀리초 대기

		}
	}()

	go func() {
		for {
			c2 <- "hello world"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case i := <-c1: //채널 c1에 값이 들어왔다면 값을 꺼내서 i에 대입
				fmt.Println("c1", i)
			case s := <-c2: //채널 c2 에 값이 들어왔다면 값을 꺼내서 s에 대입
				fmt.Println("c2", s)
				//case <-c1:
				//	fmt.Println("값 사용하지 않을때")

				//default: 케이스가 없을때 실행될 값. switch 랑 똑같음.
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
