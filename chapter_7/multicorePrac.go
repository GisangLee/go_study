package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정

	fmt.Println(runtime.GOMAXPROCS(0))

	s := "hello world"

	for i := 0; i < 100; i++ {
		go func(n int) { fmt.Println(s, n) }(i)
		fmt.Scanln()
	}
}
