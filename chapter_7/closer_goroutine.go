package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(1)

	s := "hello world"

	//클로저를 고루틴으로 실행할때 반복문 안에서 변수 사용 주의.
	//일반 클로저는 반복문 안에서 순서대로 실행되지만 고루틴 클로저는 반복문이 끝난뒤에 고루틴이 실행됨
	//의도대로 동작하게 하려면 반복문의 변수를 매개변수로 받아줘야함.
	//변화하는 변수는 꼭 매개변수로 넘겨주기
	//두가지 방식의 비교
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	// 매개변수로 안넣어주고 사용. 위와 결과값이 다르다 반복문이 끝난뒤에 고루틴이 실행됨
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println(s, i)
	//	}()
	//}

	fmt.Scanln()
}
