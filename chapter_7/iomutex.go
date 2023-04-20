package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 읽기 락 :읽기 락 끼리는 서로 막지 않음 다만 내용변경이 되면 안되기에 쓰기락은 막음
// 쓰기 락 : 읽기 쓰기 전부 막음
func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int = 0
	var rwMutex = new(sync.RWMutex) //읽기 쓰기 뮤텍스 생성

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock() //쓰기 뮤텍스 잠금.
			data += 1      // data에 값 쓰기
			fmt.Println("write:", data)
			time.Sleep(10 * time.Millisecond) //10 밀리초 대기
			rwMutex.Unlock()                  //쓰기 뮤텍스 해제
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock() //읽기 뮤텍스 잠금
			// data에 값 쓰기
			fmt.Println("read1:", data)
			time.Sleep(1 * time.Millisecond) //10 밀리초 대기
			rwMutex.RUnlock()                //쓰기 뮤텍스 해제
		}
	}()
	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock() //읽기 뮤텍스 잠금
			// data에 값 쓰기
			fmt.Println("read2:", data)
			time.Sleep(1 * time.Millisecond) //10 밀리초 대기
			rwMutex.RUnlock()                //쓰기 뮤텍스 해제
		}
	}()

	time.Sleep(10 * time.Second)
}
