package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //모든 cpu 사용

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock() // 뮤텍스 잠금 ,data 슬라이스 보호 시작
			data = append(data, 1)
			mutex.Unlock() //뮤텍스 잠금 해제 , data 슬라이스 보호 종료 ,mutex 쌍이 안맞으면 deadlock 발생
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock() // 뮤텍스 잠금 ,data 슬라이스 보호 시작
			data = append(data, 1)
			mutex.Unlock() //뮤텍스 잠금 해제 , data 슬라이스 보호 종료
			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(data)) //정확히 값 2000 나옴
}
