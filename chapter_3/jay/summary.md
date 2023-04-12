## CHAPTER 3 문자열 및 자료구조
### 3.1 문자열
바이트들의 연속
- []byte
- string: **읽기 전용**
### 3.1.1 유니코드 처리
Go 언어의 소스 코드는 UTF-8 로 되어 있음
- char: 1바이트
- byte: uint8
- rune: int32
    - 유니코드 포인트를 담음
    - UTF-8로 표현 가능한 길이는 최대 6바이트지만 다른 인코딩과의 호환을 위해 4바이트까지만 사용[^1]
```go
for i, r := range "가나다" { // i: int, r: rune
	fmt.Println(i, r)
}
fmt.Println(len("가나다")) // 9

for _, r := range "가갛힣" {
	fmt.Println(string(r), r) // rune -> string 변환
}
```

### 3.1.2 Example 테스트
테스트 방법
```sh
go test hangul_test.go hangul.go
```
hangul_test 를 빌드하는데 필요한 다른 모든 파일의 이름도 지정
- `Output:`: 표준 출력과 비교
#### Reference
- https://pkg.go.dev/testing#hdr-Examples
### 3.1.3 바이트 단위 처리
- [hangul_test.go](./hangul_test.go)
- string: 어떤 문자들이 들어 있는지를 중시
- []byte: 실제 바이트 표현이 어떤지를 중시
### 3.1.4 패키지 문서
- 함수 주석: 함수 이름으로 시작
- 패키지 주석: "Package" 로 시작
### 3.1.5 문자열 잇기
- [hangul_test.go](./hangul_test.go)
### 3.1.6 문자열을 숫자로
문자열 -> 숫자
- strconv.Atoi
- strconv.ParseInt
- fmt.Sscanf

숫자 -> 문자열
- strconv.Itoa
- strconv.FormatInt
- fmt.Sprintf
### 3.2 배열과 슬라이스
### 3.2.1 배열
### 3.2.2 슬라이스
길이와 용량을 갖고 있고 길이가 변할 수 있는 구조
### 3.2.3 슬라이스 덧붙이기
- append()
    - 가변 인자를 받음
- `...`
    - 슬라이스의 값들을 인자로 늘어 놓을 수 있음
### 3.2.4 슬라이스 용량
- 슬라이스의 최대 길이
    - 해당 용량까지는 새로 복사가 이루어지지 않음
    - 길이까지만 배열을 초기화
#### Reference
- 슬라이스의 길이와 용량 차이 관련 글: https://teivah.medium.com/slice-length-vs-capacity-in-go-af71a754b7d8
- 슬라이스의 메모리 누수 관련 글: https://hsleep.medium.com/memory-leak-in-golang-slice-41f93f800af
### 3.2.5 슬라이스의 내부 구현
- 배열을 가리키는 포인터
- 길이
- 용량
### 3.2.6 슬라이스 복사
- [slice_test.go](./slice_test.go)
```go
// 1. copy 함수 사용
dest := make([]int, len(src))
copy(dest, src) // len(src) 와 len(desc) 중에서 더 작은 값만큼만 복사, 실제로 몇 개가 복사되었는지 반환
// 2. append 함수 사용
desc := append([]int(nil), src...)
```
### 3.2.7 슬라이스 삽입 및 삭제
```go
// 삽입
a = append(a[:i+1], a[i:]...)
a[i] = x
// 삭제
a = append(a[:i], a[i+1]...)
```
- 삭제되는 슬라이스 내부에 포인터가 있는 경우 메모리 누수가 일어남
    ```go
    copy(a[i:], a[i+k])
    for i := 0; i < k; i++ {
        a[len(a)-1-i] = nil // 생략시 메모리 누수 위험
    }
    a = a[:len(a)-k]
    ```
### 3.2.8 스택
- [calculator.go](./calculator.go)
### 3.3 맵
해시테이블로 구현
### 3.3.1 맵 사용하기
- [map_test.go](./map_test.go)
- 순서가 없음
- 삭제
  - `delete(m, key)`
### 3.3.2 집합
- `map[rune]struct{}{}` 사용
    - 빈 구조체를 사용해 메모리 낭비가 없음
### 3.3.3 맵의 한계
- 스레드 안전하지 않음
- 키에는 불변값만 가능
### 3.4 입출력
표준 라이브러리 io 에서 제공
### 3.4.1 io.Reader와 io.Writer
바이트들을 읽고 쓸 수 있는 인터페이스
### 3.4.2 파일 읽기
```go
package fmt

func Fscanf(r io.Reader, format string, a ...any) (n int, err err)
```
```go
f, err := os.Open(fiename)
if err != nil {
	return err
}
defer f.Close()
var num int
if _, err := fmt.Fscanf(f, "%d\n", &num); err == nil {
	// 읽은 num 값 사용
}
```
### 3.4.3 파일 쓰기
```go
package fmt

func Fprintf(w io.Writer, format stirng, a ...any) (n int, err error)
```
```go
f, err := os.Create(filename)
if err != nil {
	return err
}
defer f.Close()
if _, err := fmt.Fprintf(f, "%d\n", num); err != nil {
	retrun err
}
```
### 3.4.4 텍스트 리스트 읽고 쓰기
- [io.go](./io.go)
### 3.4.5 그래프의 인접 리스트 읽고 쓰기
### 3.5 마치며
### 3.6 연습문제 

---
[^1]: https://namu.wiki/w/UTF-8