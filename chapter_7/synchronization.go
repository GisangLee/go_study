package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU()) //모든 cpu 사용

	var data = []int{} //int 슬라이스 생성

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, i) //data 슬라이스에 i 추가
			runtime.Gosched()      //다른 고루틴이 cpu 사용할 수 있도록 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)
			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(len(data))

	// 이 경우 결과값이 2000이 나와야 하는데 1800~1990 사이의 값이 나옴.
	//두 고루틴이 경합을 벌이면서 동시접근해서 함수가 정확히 처리되지 않은것
	//이러한 상황을 Race condition 이라고 함
	//cpu 코어가 1개면 경쟁 조건이 발생 안할수도 있음. runtime.GOMAXPROCS(1) 의 상황
}
