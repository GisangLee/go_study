package jay

import "fmt"

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("Go 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))
	// Output:
	// false
	// true
	// false
}

/*
*
3.1.3 바이트 단위 처리
*/
func Example_printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ { // for range 문이 아니라면 byte 단위로 동작
		fmt.Printf("%x:", s[i]) // %x: 16진수 숫자 형식
	}
	fmt.Println()
	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printBytes2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s) // % x: 바이트 단위로 한 칸씩 띄움
	// Output:
	// eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b)) // []byte -> string 형변환
	// Output:
	// 각나다
}

/*
*
3.1.5 문자열 잇기
*/
func Example_strCat() {
	s := "abc"
	ps := &s   // ps가 가리키고 있는 값도 변경
	s += "def" // fmt.Sprint(), fmt.Sprintf(), strings.Join() 으로 대체 가능
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}
