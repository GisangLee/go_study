package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("hi", i)
		}
		//채널을 닫음
		close(c)
	}()

	for i := range c { //range 사용해서 채널에서 값을 꺼냄. rnage 사용 시 채널이 닫힐때까지 반복하면서 값을 꺼냄. //range 는 값이 몇개인지 모르는 상황에서 유용
		fmt.Println(i)
	}
}
